package services

import (
	"context"
	"errors"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ariefsn/gotik/constant"
	"github.com/ariefsn/gotik/entities"
	"github.com/ariefsn/gotik/helper"
	"github.com/ariefsn/gotik/logger"
	"github.com/nats-io/nats.go"
	"github.com/playwright-community/playwright-go"
)

type TiktokService struct {
	nc   *nats.Conn
	repo entities.TiktokRepository
}

// GetQueueStatus implements entities.TiktokService.
func (t *TiktokService) GetQueueStatus(ctx context.Context, keyword string) (entities.TiktokVideoStats, error) {
	return t.repo.GetQueueStatus(ctx, keyword)
}

// SetQueueStatus implements entities.TiktokService.
func (t *TiktokService) SetQueueStatus(ctx context.Context, keyword string, status constant.TiktokQueueStatus) error {
	res, _ := t.repo.GetQueueStatus(ctx, keyword)
	res.Stats = status
	return t.repo.SetQueueStatus(ctx, keyword, res)
}

// GetById implements entities.TiktokService.
func (t *TiktokService) GetById(ctx context.Context, id string) (*entities.VideoItem, error) {
	return t.repo.GetById(ctx, id)
}

// GetNotif implements entities.TiktokService.
func (t *TiktokService) GetNotif(ctx context.Context, keyword string) ([]*entities.VideoItem, error) {
	return t.repo.GetNotif(ctx, keyword)
}

// SendQueue implements entities.TiktokService.
func (t *TiktokService) SendQueue(ctx context.Context, keyword string) {
	msg := entities.QueTiktokSearch{
		Keyword: keyword,
	}
	err := t.nc.Publish(string(constant.TiktokSubscribeKey), helper.ToBytes(msg))
	if err != nil {
		logger.Error(err)
	}
}

// Get implements entities.TiktokService.
func (t *TiktokService) Get(ctx context.Context, keyword string) ([]*entities.VideoItem, entities.TiktokVideoStats, error) {
	status, err := t.repo.GetQueueStatus(ctx, keyword)

	if status.Stats == constant.TiktokQueueStatusNone {
		return []*entities.VideoItem{}, status, err
	}

	// Check if token expired
	if time.Now().Unix() > status.Expiry {
		status.Stats = constant.TiktokQueueStatusNone
		t.repo.Store(ctx, keyword, []*entities.VideoItem{})
		return []*entities.VideoItem{}, status, errors.New("token expired")
	}

	res, err := t.repo.Search(ctx, keyword)

	return res, status, err
}

// Search implements entities.TiktokService.
func (t *TiktokService) Search(ctx context.Context, keyword string) ([]*entities.VideoItem, error) {
	existing, _ := t.repo.Search(ctx, keyword)
	if len(existing) > 0 {
		logger.Info("found existing data, not needs to scrape")
		return existing, nil
	}

	logger.Info("scraping data for keyword = " + keyword)
	pw, err := playwright.Run()
	if err != nil {
		log.Fatal(err)
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatal(err)
	}

	type IResponse struct {
		IsFinished bool `json:"isFinished"`
		Response   playwright.Response
		Data       []*entities.VideoItem
		Offset     int
	}

	responseChan := make(chan IResponse)

	page.OnResponse(func(r playwright.Response) {
		url := r.URL()
		if strings.Contains(url, "https://www.tiktok.com/api/search/general/full") {
			go func() {
				var res entities.TiktokSearchResponse
				r.JSON(&res)
				cookies, _ := page.Context().Cookies()
				token := ""
				tokenExpiry := int64(0)

				for _, c := range cookies {
					if strings.ToLower(c.Name) == "tt_chain_token" {
						token = c.Value
						tokenExpiry = int64(c.Expires)
						break
					}
				}

				var results []*entities.VideoItem

				if token == "" {
					logger.Info("tt_chain_token not found")
				} else {
					logger.Info("tt_chain_token found " + token)
					stats, _ := t.repo.GetQueueStatus(ctx, keyword)
					stats.Token = token
					stats.Expiry = tokenExpiry
					t.repo.SetQueueStatus(ctx, keyword, stats)
					for _, v := range res.Data {
						if v.Item.Video.PlayAddr == "" {
							continue
						}
						v.Item.Video.SetPlayAddrPublic(token)
						results = append(results, v)
					}
				}

				chanResult := IResponse{
					Response: r,
					Data:     results,
					Offset:   res.Cursor / 12,
				}

				if len(results) == 0 {
					logger.Info("finished scraping data for keyword = " + keyword + " because no results found")
					stats, _ := t.repo.GetQueueStatus(ctx, keyword)
					stats.Stats = constant.TiktokQueueStatusFinished
					t.repo.SetQueueStatus(ctx, keyword, stats)
					return
				}

				sort.Slice(results, func(i, j int) bool {
					return results[i].Item.CreateTime < results[j].Item.CreateTime
				})

				lastItem := results[len(results)-1]
				unixCreatedTime := lastItem.Item.CreateTime
				// if createdTime is less than 7 day, then scroll
				sevenDaysAgo := time.Now().AddDate(0, 0, -7)
				timeDiff := int64(unixCreatedTime) - sevenDaysAgo.Unix()
				if (timeDiff < 0 || res.Cursor < 50) && res.HasMore == 1 {
					chanResult.IsFinished = false
					responseChan <- chanResult
				} else {
					logger.Info("finished scraping data for keyword = " + keyword)
					stats, _ := t.repo.GetQueueStatus(ctx, keyword)
					stats.Stats = constant.TiktokQueueStatusFinished
					t.repo.SetQueueStatus(ctx, keyword, stats)
					close(responseChan)
				}
			}()
		}
	})

	_, err = page.Goto("https://www.tiktok.com/search?q="+keyword, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	})

	if err != nil {
		log.Fatal(err)
	}

	stats, _ := t.repo.GetQueueStatus(ctx, keyword)
	stats.Stats = constant.TiktokQueueStatusInitiated
	t.repo.SetQueueStatus(ctx, keyword, stats)

	for msg := range responseChan {
		logger.Info("storing result data for keyword = " + keyword)
		t.repo.Store(ctx, keyword, msg.Data)

		if !msg.IsFinished {
			// Scroll
			logger.Info("scrolling data for keyword = " + keyword + " offset = " + strconv.Itoa(msg.Offset))
			page.Mouse().Wheel(0, float64(msg.Offset*page.ViewportSize().Height))
		}
	}

	var results []*entities.VideoItem
	return results, nil
}

func NewTiktokService(repo entities.TiktokRepository, nc *nats.Conn) entities.TiktokService {
	return &TiktokService{
		nc:   nc,
		repo: repo,
	}
}
