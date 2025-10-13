package cachelibs

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_cacheImpl_Get(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		store        map[string]cacheItem[string]
		key          string
		assertions   assert.ErrorAssertionFunc
		want         string
		wantErrMsg   string
		checkDeleted bool
	}{
		{
			name: "key exists and not expired",
			store: map[string]cacheItem[string]{
				"key1": {value: "value1", expiresAt: time.Now().Add(1 * time.Hour)},
			},
			key:        "key1",
			assertions: assert.NoError,
			want:       "value1",
		},
		{
			name: "key exists but expired",
			store: map[string]cacheItem[string]{
				"key2": {value: "value2", expiresAt: time.Now().Add(-1 * time.Hour)},
			},
			key:          "key2",
			assertions:   assert.Error,
			want:         "",
			wantErrMsg:   "data not found or expired",
			checkDeleted: true,
		},
		{
			name: "key does not exist",
			store: map[string]cacheItem[string]{
				"key3": {value: "value3", expiresAt: time.Now().Add(1 * time.Hour)},
			},
			key:        "key4",
			assertions: assert.Error,
			want:       "",
			wantErrMsg: "data not found or expired",
		},
		{
			name: "key exists with infinite TTL",
			store: map[string]cacheItem[string]{
				"key5": {value: "value5", expiresAt: time.Time{}},
			},
			key:        "key5",
			assertions: assert.NoError,
			want:       "value5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cache := &cacheImpl[string]{store: tt.store}
			got, err := cache.Get(tt.key)
			tt.assertions(t, err)
			if err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
			assert.Equal(t, tt.want, got)
			if tt.checkDeleted {
				_, exists := cache.store[tt.key]
				assert.False(t, exists)
			}
		})
	}
}

func Test_cacheImpl_Set(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		opts         CacheOptions
		initialStore map[string]cacheItem[string]
		key          string
		value        string
		assertions   assert.ErrorAssertionFunc
		wantErrMsg   string
	}{
		{
			name:         "set value with unlimited size",
			opts:         CacheOptions{Ttl: 1 * time.Hour, MaxSize: 0},
			initialStore: map[string]cacheItem[string]{},
			key:          "key1",
			value:        "value1",
			assertions:   assert.NoError,
		},
		{
			name:         "set value with size limit not reached",
			opts:         CacheOptions{Ttl: 1 * time.Hour, MaxSize: 2},
			initialStore: map[string]cacheItem[string]{},
			key:          "key2",
			value:        "value2",
			assertions:   assert.NoError,
		},
		{
			name: "set value with size limit reached",
			opts: CacheOptions{Ttl: 1 * time.Hour, MaxSize: 1},
			initialStore: map[string]cacheItem[string]{
				"key1": {value: "value1", expiresAt: time.Now().Add(1 * time.Hour)},
			},
			key:        "key3",
			value:      "value3",
			assertions: assert.Error,
			wantErrMsg: "cache size limit reached",
		},
		{
			name:         "set value with infinite TTL",
			opts:         CacheOptions{Ttl: 0, MaxSize: 0},
			initialStore: map[string]cacheItem[string]{},
			key:          "key4",
			value:        "value4",
			assertions:   assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cache := &cacheImpl[string]{store: tt.initialStore, opts: tt.opts}
			err := cache.Set(tt.key, tt.value)
			tt.assertions(t, err)
			if err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
			if err == nil {
				item, exists := cache.store[tt.key]
				assert.True(t, exists)
				assert.Equal(t, tt.value, item.value)
				if tt.opts.Ttl > 0 {
					assert.WithinDuration(t, time.Now().Add(tt.opts.Ttl), item.expiresAt, time.Second)
				} else {
					assert.True(t, item.expiresAt.IsZero())
				}
			}
		})
	}
}

func TestNewCache(t *testing.T) {
	t.Parallel()
	opts := CacheOptions{Ttl: 30 * time.Minute, MaxSize: 100, RefreshCycle: 5 * time.Minute}
	cache := NewCache[string](context.Background(), opts)

	assert.NotNil(t, cache)
	assert.Equal(t, opts, cache.(*cacheImpl[string]).opts)
	assert.Equal(t, map[string]cacheItem[string]{}, cache.(*cacheImpl[string]).store)
}

func Test_cacheImpl_Refresh(t *testing.T) {
	t.Parallel()
	store := map[string]cacheItem[string]{
		"key1": {value: "value1", expiresAt: time.Now().Add(-1 * time.Hour)}, // expired
		"key2": {value: "value2", expiresAt: time.Now().Add(1 * time.Hour)},  // not expired
		"key3": {value: "value3", expiresAt: time.Time{}},                    // infinite TTL
	}
	cache := &cacheImpl[string]{store: store}

	cache.Refresh()

	_, exists1 := cache.store["key1"]
	_, exists2 := cache.store["key2"]
	_, exists3 := cache.store["key3"]

	assert.False(t, exists1)
	assert.True(t, exists2)
	assert.True(t, exists3)
}
