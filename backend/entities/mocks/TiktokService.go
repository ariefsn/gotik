// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	constant "github.com/ariefsn/gotik/constant"

	entities "github.com/ariefsn/gotik/entities"

	mock "github.com/stretchr/testify/mock"
)

// TiktokService is an autogenerated mock type for the TiktokService type
type TiktokService struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, keyword
func (_m *TiktokService) Get(ctx context.Context, keyword string) ([]*entities.VideoItem, entities.TiktokVideoStats, error) {
	ret := _m.Called(ctx, keyword)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []*entities.VideoItem
	var r1 entities.TiktokVideoStats
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*entities.VideoItem, entities.TiktokVideoStats, error)); ok {
		return rf(ctx, keyword)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entities.VideoItem); ok {
		r0 = rf(ctx, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.VideoItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) entities.TiktokVideoStats); ok {
		r1 = rf(ctx, keyword)
	} else {
		r1 = ret.Get(1).(entities.TiktokVideoStats)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, keyword)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetById provides a mock function with given fields: ctx, id
func (_m *TiktokService) GetById(ctx context.Context, id string) (*entities.VideoItem, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *entities.VideoItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entities.VideoItem, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entities.VideoItem); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.VideoItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotif provides a mock function with given fields: ctx, keyword
func (_m *TiktokService) GetNotif(ctx context.Context, keyword string) ([]*entities.VideoItem, error) {
	ret := _m.Called(ctx, keyword)

	if len(ret) == 0 {
		panic("no return value specified for GetNotif")
	}

	var r0 []*entities.VideoItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*entities.VideoItem, error)); ok {
		return rf(ctx, keyword)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entities.VideoItem); ok {
		r0 = rf(ctx, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.VideoItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQueueStatus provides a mock function with given fields: ctx, keyword
func (_m *TiktokService) GetQueueStatus(ctx context.Context, keyword string) (entities.TiktokVideoStats, error) {
	ret := _m.Called(ctx, keyword)

	if len(ret) == 0 {
		panic("no return value specified for GetQueueStatus")
	}

	var r0 entities.TiktokVideoStats
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entities.TiktokVideoStats, error)); ok {
		return rf(ctx, keyword)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entities.TiktokVideoStats); ok {
		r0 = rf(ctx, keyword)
	} else {
		r0 = ret.Get(0).(entities.TiktokVideoStats)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, keyword
func (_m *TiktokService) Search(ctx context.Context, keyword string) ([]*entities.VideoItem, error) {
	ret := _m.Called(ctx, keyword)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 []*entities.VideoItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*entities.VideoItem, error)); ok {
		return rf(ctx, keyword)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*entities.VideoItem); ok {
		r0 = rf(ctx, keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.VideoItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendQueue provides a mock function with given fields: ctx, keyword
func (_m *TiktokService) SendQueue(ctx context.Context, keyword string) {
	_m.Called(ctx, keyword)
}

// SetQueueStatus provides a mock function with given fields: ctx, keyword, status
func (_m *TiktokService) SetQueueStatus(ctx context.Context, keyword string, status constant.TiktokQueueStatus) error {
	ret := _m.Called(ctx, keyword, status)

	if len(ret) == 0 {
		panic("no return value specified for SetQueueStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, constant.TiktokQueueStatus) error); ok {
		r0 = rf(ctx, keyword, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTiktokService creates a new instance of TiktokService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTiktokService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TiktokService {
	mock := &TiktokService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}