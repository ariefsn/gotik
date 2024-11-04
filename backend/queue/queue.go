package queue

import (
	"context"
	"fmt"

	"github.com/ariefsn/gotik/constant"
	"github.com/ariefsn/gotik/entities"
	"github.com/ariefsn/gotik/helper"
	"github.com/ariefsn/gotik/logger"
	"github.com/nats-io/nats.go"
)

type Queue struct {
	c      *nats.Conn
	tiktok entities.TiktokService
}

type QueueOptions struct {
	Conn          *nats.Conn
	TiktokService entities.TiktokService
}

func NewQueue(opt QueueOptions) *Queue {
	return &Queue{
		c:      opt.Conn,
		tiktok: opt.TiktokService,
	}
}

func (q *Queue) Listen() {
	queueGroup := "gotik"

	q.c.QueueSubscribe(string(constant.TiktokSubscribeKey), queueGroup, func(m *nats.Msg) {
		logger.Info(fmt.Sprintf("tiktok payload received: %s", string(m.Data)))

		ctx := context.Background()

		payload, err := helper.FromBytes[entities.QueTiktokSearch](m.Data)
		if err != nil {
			logger.Error(fmt.Errorf("failed to parse tiktok payload: %w", err))
		} else {
			q.tiktok.Search(ctx, payload.Keyword)
		}
	})
}
