package repository_test

import (
	"context"
	"testing"

	"github.com/ariefsn/gotik/apps/tiktok/repository"
	"github.com/ariefsn/gotik/constant"
	"github.com/ariefsn/gotik/entities"
	"github.com/stretchr/testify/assert"
)

func TestSetQueueStatus(t *testing.T) {
	repo := repository.NewTiktokRepository()

	err := repo.SetQueueStatus(context.Background(), "test", entities.TiktokVideoStats{
		Stats: constant.TiktokQueueStatusNone,
	})

	assert.Nil(t, err)
}

func TestGetQueueStatus(t *testing.T) {
	repo := repository.NewTiktokRepository()

	_, err := repo.GetQueueStatus(context.Background(), "test")

	assert.Nil(t, err)

	res, _ := repo.GetQueueStatus(context.Background(), "test2")

	assert.Equal(t, constant.TiktokQueueStatusNone, res.Stats)
}

func TestStore(t *testing.T) {
	repo := repository.NewTiktokRepository()

	repo.Store(context.Background(), "test", []*entities.VideoItem{
		{
			Common: entities.Common{
				DocIDStr: "001",
			},
			Item: entities.Item{
				ID: "001",
			},
		},
	})
}

func TestGetNotif(t *testing.T) {
	repo := repository.NewTiktokRepository()

	res, err := repo.GetNotif(context.Background(), "test")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestSearch(t *testing.T) {
	repo := repository.NewTiktokRepository()

	res, err := repo.Search(context.Background(), "test")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGetById(t *testing.T) {
	repo := repository.NewTiktokRepository()

	res, err := repo.GetById(context.Background(), "001")

	assert.Nil(t, err)
	assert.NotNil(t, res)

	res, err = repo.GetById(context.Background(), "002")

	assert.Nil(t, res)
	assert.NotNil(t, err)

}
