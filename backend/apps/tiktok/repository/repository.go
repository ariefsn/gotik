package repository

import (
	"context"
	"errors"

	"github.com/ariefsn/gotik/constant"
	"github.com/ariefsn/gotik/entities"
)

var videoStorage = map[string][]*entities.VideoItem{}

var videoNotifStorage = map[string][]*entities.VideoItem{}

var videoStats = map[string]*entities.TiktokVideoStats{}

type TiktokRepository struct {
}

// SetQueueStatus implements entities.TiktokRepository.
func (t *TiktokRepository) SetQueueStatus(ctx context.Context, keyword string, stats entities.TiktokVideoStats) error {
	videoStats[keyword] = &stats
	return nil
}

// GetQueueStatus implements entities.TiktokRepository.
func (t *TiktokRepository) GetQueueStatus(ctx context.Context, keyword string) (entities.TiktokVideoStats, error) {
	if val, ok := videoStats[keyword]; ok {
		return *val, nil
	}

	return entities.TiktokVideoStats{
		Stats: constant.TiktokQueueStatusNone,
	}, nil
}

// GetById implements entities.TiktokRepository.
func (t *TiktokRepository) GetById(ctx context.Context, id string) (*entities.VideoItem, error) {
	for _, videos := range videoStorage {
		for _, v := range videos {
			if v.Item.ID == id {
				return v, nil
			}
		}
	}

	return nil, errors.New("video not found")
}

// GetNotif implements entities.TiktokRepository.
func (t *TiktokRepository) GetNotif(ctx context.Context, keyword string) ([]*entities.VideoItem, error) {
	return videoNotifStorage[keyword], nil
}

// Search implements entities.TiktokRepository.
func (t *TiktokRepository) Search(ctx context.Context, keyword string) ([]*entities.VideoItem, error) {
	return videoStorage[keyword], nil
}

// Store implements entities.TiktokRepository.
func (t *TiktokRepository) Store(ctx context.Context, keyword string, videos []*entities.VideoItem) {
	tmp := videoStorage[keyword]
	tmp = append(tmp, videos...)
	videoStorage[keyword] = tmp
	videoNotifStorage[keyword] = videos
}

func NewTiktokRepository() entities.TiktokRepository {
	return &TiktokRepository{}
}
