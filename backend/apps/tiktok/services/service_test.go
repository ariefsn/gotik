package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ariefsn/gotik/apps/tiktok/services"
	"github.com/ariefsn/gotik/constant"
	"github.com/ariefsn/gotik/entities"
	"github.com/ariefsn/gotik/entities/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetQueueStatus(t *testing.T) {
	repo := mocks.NewTiktokRepository(t)
	repo.On("GetQueueStatus", mock.Anything, mock.Anything).Return(entities.TiktokVideoStats{
		Stats: constant.TiktokQueueStatusNone,
	}, nil)

	service := services.NewTiktokService(repo, nil)
	res, err := service.GetQueueStatus(context.Background(), "test")

	assert.Equal(t, constant.TiktokQueueStatusNone, res.Stats)
	assert.Nil(t, err)
}

func TestSetQueueStatus(t *testing.T) {
	repo := mocks.NewTiktokRepository(t)
	repo.On("GetQueueStatus", mock.Anything, mock.Anything).Return(entities.TiktokVideoStats{
		Stats: constant.TiktokQueueStatusNone,
	}, nil)
	repo.On("SetQueueStatus", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	service := services.NewTiktokService(repo, nil)
	err := service.SetQueueStatus(context.Background(), "test", constant.TiktokQueueStatusNone)

	assert.Nil(t, err)
}

func TestGetById(t *testing.T) {
	repo := mocks.NewTiktokRepository(t)
	repo.On("GetById", mock.Anything, mock.Anything).Return(nil, errors.New("data not found"))

	service := services.NewTiktokService(repo, nil)
	res, err := service.GetById(context.Background(), "test")

	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestGetNotif(t *testing.T) {
	repo := mocks.NewTiktokRepository(t)
	repo.On("GetNotif", mock.Anything, mock.Anything).Return([]*entities.VideoItem{}, nil)

	service := services.NewTiktokService(repo, nil)
	res, err := service.GetNotif(context.Background(), "test")

	assert.Equal(t, 0, len(res))
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	repo := mocks.NewTiktokRepository(t)
	repo.On("GetQueueStatus", mock.Anything, mock.Anything).Return(entities.TiktokVideoStats{
		Stats: constant.TiktokQueueStatusNone,
	}, nil)
	// repo.On("Store", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	// repo.On("Search", mock.Anything, mock.Anything).Return([]*entities.VideoItem{}, nil)

	service := services.NewTiktokService(repo, nil)
	res, status, err := service.Get(context.Background(), "test")

	assert.Equal(t, constant.TiktokQueueStatusNone, status.Stats)
	assert.Equal(t, 0, len(res))
	assert.Nil(t, err)
}
