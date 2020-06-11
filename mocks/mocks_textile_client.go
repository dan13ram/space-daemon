// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import (
	client "github.com/FleekHQ/space-poc/core/textile/client"
	apiclient "github.com/textileio/go-threads/api/client"

	context "context"

	io "io"

	mock "github.com/stretchr/testify/mock"

	path "github.com/ipfs/interface-go-ipfs-core/path"

	thread "github.com/textileio/go-threads/core/thread"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateBucket provides a mock function with given fields: ctx, bucketSlug
func (_m *Client) CreateBucket(ctx context.Context, bucketSlug string) (*client.TextileBucketRoot, error) {
	ret := _m.Called(ctx, bucketSlug)

	var r0 *client.TextileBucketRoot
	if rf, ok := ret.Get(0).(func(context.Context, string) *client.TextileBucketRoot); ok {
		r0 = rf(ctx, bucketSlug)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.TextileBucketRoot)
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

// CreateDirectory provides a mock function with given fields: ctx, bucketKey, _a2
func (_m *Client) CreateDirectory(ctx context.Context, bucketKey string, _a2 string) (path.Resolved, path.Path, error) {
	ret := _m.Called(ctx, bucketKey, _a2)

	var r0 path.Resolved
	if rf, ok := ret.Get(0).(func(context.Context, string, string) path.Resolved); ok {
		r0 = rf(ctx, bucketKey, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(path.Resolved)
		}
	}

	var r1 path.Path
	if rf, ok := ret.Get(1).(func(context.Context, string, string) path.Path); ok {
		r1 = rf(ctx, bucketKey, _a2)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(path.Path)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, bucketKey, _a2)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteDirOrFile provides a mock function with given fields: ctx, bucketKey, _a2
func (_m *Client) DeleteDirOrFile(ctx context.Context, bucketKey string, _a2 string) (path.Resolved, error) {
	ret := _m.Called(ctx, bucketKey, _a2)

	var r0 path.Resolved
	if rf, ok := ret.Get(0).(func(context.Context, string, string) path.Resolved); ok {
		r0 = rf(ctx, bucketKey, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(path.Resolved)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, bucketKey, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileExists provides a mock function with given fields: ctx, key, _a2, r
func (_m *Client) FileExists(ctx context.Context, key string, _a2 string, r io.Reader) (bool, error) {
	ret := _m.Called(ctx, key, _a2, r)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader) bool); ok {
		r0 = rf(ctx, key, _a2, r)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, io.Reader) error); ok {
		r1 = rf(ctx, key, _a2, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FolderExists provides a mock function with given fields: ctx, key, _a2
func (_m *Client) FolderExists(ctx context.Context, key string, _a2 string) (bool, error) {
	ret := _m.Called(ctx, key, _a2)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, key, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, _a2)
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

// GetFile provides a mock function with given fields: ctx, bucketKey, _a2, w
func (_m *Client) GetFile(ctx context.Context, bucketKey string, _a2 string, w io.Writer) error {
	ret := _m.Called(ctx, bucketKey, _a2, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Writer) error); ok {
		r0 = rf(ctx, bucketKey, _a2, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetThreadsConnection provides a mock function with given fields:
func (_m *Client) GetThreadsConnection() (*apiclient.Client, error) {
	ret := _m.Called()

	var r0 *apiclient.Client
	if rf, ok := ret.Get(0).(func() *apiclient.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiclient.Client)
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

// ListBuckets provides a mock function with given fields: ctx
func (_m *Client) ListBuckets(ctx context.Context) ([]*client.TextileBucketRoot, error) {
	ret := _m.Called(ctx)

	var r0 []*client.TextileBucketRoot
	if rf, ok := ret.Get(0).(func(context.Context) []*client.TextileBucketRoot); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*client.TextileBucketRoot)
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

// ListDirectory provides a mock function with given fields: ctx, bucketKey, _a2
func (_m *Client) ListDirectory(ctx context.Context, bucketKey string, _a2 string) (*client.TextileDirEntries, error) {
	ret := _m.Called(ctx, bucketKey, _a2)

	var r0 *client.TextileDirEntries
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *client.TextileDirEntries); ok {
		r0 = rf(ctx, bucketKey, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.TextileDirEntries)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, bucketKey, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *Client) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartAndBootstrap provides a mock function with given fields: ctx
func (_m *Client) StartAndBootstrap(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Client) Stop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadFile provides a mock function with given fields: ctx, bucketKey, _a2, reader
func (_m *Client) UploadFile(ctx context.Context, bucketKey string, _a2 string, reader io.Reader) (path.Resolved, path.Path, error) {
	ret := _m.Called(ctx, bucketKey, _a2, reader)

	var r0 path.Resolved
	if rf, ok := ret.Get(0).(func(context.Context, string, string, io.Reader) path.Resolved); ok {
		r0 = rf(ctx, bucketKey, _a2, reader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(path.Resolved)
		}
	}

	var r1 path.Path
	if rf, ok := ret.Get(1).(func(context.Context, string, string, io.Reader) path.Path); ok {
		r1 = rf(ctx, bucketKey, _a2, reader)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(path.Path)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, io.Reader) error); ok {
		r2 = rf(ctx, bucketKey, _a2, reader)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
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