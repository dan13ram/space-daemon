// Code generated by mockery v2.0.3. DO NOT EDIT.

package mocks

import (
	config "github.com/FleekHQ/space-daemon/config"
	client "github.com/textileio/go-threads/api/client"

	context "context"

	mock "github.com/stretchr/testify/mock"

	textile "github.com/FleekHQ/space-daemon/core/textile"

	thread "github.com/textileio/go-threads/core/thread"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateBucket provides a mock function with given fields: ctx, bucketSlug
func (_m *Client) CreateBucket(ctx context.Context, bucketSlug string) (textile.Bucket, error) {
	ret := _m.Called(ctx, bucketSlug)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context, string) textile.Bucket); ok {
		r0 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, bucketSlug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBaseThreadsContext provides a mock function with given fields: ctx
func (_m *Client) GetBaseThreadsContext(ctx context.Context) (context.Context, error) {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucket provides a mock function with given fields: ctx, slug
func (_m *Client) GetBucket(ctx context.Context, slug string) (textile.Bucket, error) {
	ret := _m.Called(ctx, slug)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context, string) textile.Bucket); ok {
		r0 = rf(ctx, slug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketContext provides a mock function with given fields: ctx, bucketSlug
func (_m *Client) GetBucketContext(ctx context.Context, bucketSlug string) (context.Context, *thread.ID, error) {
	ret := _m.Called(ctx, bucketSlug)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, string) context.Context); ok {
		r0 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 *thread.ID
	if rf, ok := ret.Get(1).(func(context.Context, string) *thread.ID); ok {
		r1 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*thread.ID)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, bucketSlug)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDefaultBucket provides a mock function with given fields: ctx
func (_m *Client) GetDefaultBucket(ctx context.Context) (textile.Bucket, error) {
	ret := _m.Called(ctx)

	var r0 textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context) textile.Bucket); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLocalBucketContext provides a mock function with given fields: ctx, bucketSlug
func (_m *Client) GetLocalBucketContext(ctx context.Context, bucketSlug string) (context.Context, *thread.ID, error) {
	ret := _m.Called(ctx, bucketSlug)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context, string) context.Context); ok {
		r0 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	var r1 *thread.ID
	if rf, ok := ret.Get(1).(func(context.Context, string) *thread.ID); ok {
		r1 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*thread.ID)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, bucketSlug)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetThreadsConnection provides a mock function with given fields:
func (_m *Client) GetThreadsConnection() (*client.Client, error) {
	ret := _m.Called()

	var r0 *client.Client
	if rf, ok := ret.Get(0).(func() *client.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsRunning provides a mock function with given fields:
func (_m *Client) IsRunning() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ListBuckets provides a mock function with given fields: ctx
func (_m *Client) ListBuckets(ctx context.Context) ([]textile.Bucket, error) {
	ret := _m.Called(ctx)

	var r0 []textile.Bucket
	if rf, ok := ret.Get(0).(func(context.Context) []textile.Bucket); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]textile.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Shutdown provides a mock function with given fields:
func (_m *Client) Shutdown() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartAndBootstrap provides a mock function with given fields: ctx, cfg
func (_m *Client) StartAndBootstrap(ctx context.Context, cfg config.Config) error {
	ret := _m.Called(ctx, cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.Config) error); ok {
		r0 = rf(ctx, cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitForReady provides a mock function with given fields:
func (_m *Client) WaitForReady() chan bool {
	ret := _m.Called()

	var r0 chan bool
	if rf, ok := ret.Get(0).(func() chan bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan bool)
		}
	}

	return r0
}
