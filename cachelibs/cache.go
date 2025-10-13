package cachelibs

import (
	"context"
	"errors"
	"time"
)

type Cache[T any] interface {
	Get(key string) (T, error)
	Set(key string, value T) error
	Delete(key string)
}

type cacheItem[T any] struct {
	value     T
	expiresAt time.Time // UTC
}

type CacheOptions struct {
	Ttl          time.Duration // default: infinite
	MaxSize      uint          // default: unlimited
	RefreshCycle time.Duration // delete expired data in this cycle. default: 1 minute
}

type cacheImpl[T any] struct {
	store map[string]cacheItem[T]
	opts  CacheOptions
}

func (c *cacheImpl[T]) Get(key string) (T, error) {
	if item, exists := c.store[key]; exists {
		if item.expiresAt.IsZero() || time.Now().UTC().Before(item.expiresAt) {
			return item.value, nil
		}
		c.Delete(key)
	}
	var zero T
	return zero, errors.New("data not found or expired")
}

func (c *cacheImpl[T]) Set(key string, value T) error {
	if c.opts.MaxSize > 0 && uint(len(c.store)) >= c.opts.MaxSize {
		return errors.New("cache size limit reached")
	}

	expiresAt := time.Time{}
	if c.opts.Ttl > 0 {
		expiresAt = time.Now().UTC().Add(c.opts.Ttl)
	}
	c.store[key] = cacheItem[T]{value: value, expiresAt: expiresAt}
	return nil
}

func (c *cacheImpl[T]) Delete(key string) {
	delete(c.store, key)
}

func (c *cacheImpl[T]) Refresh() {
	now := time.Now().UTC()
	for key, item := range c.store {
		if !item.expiresAt.IsZero() && now.After(item.expiresAt) {
			c.Delete(key)
		}
	}
}

func (c *cacheImpl[T]) startAutoRefresh(ctx context.Context) {
	ticker := time.NewTicker(c.opts.RefreshCycle)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.Refresh()
		case <-ctx.Done():
			return
		}
	}
}

func NewCache[T any](ctx context.Context, opts CacheOptions) Cache[T] {
	if opts.RefreshCycle <= 0 {
		opts.RefreshCycle = time.Minute
	}
	c := &cacheImpl[T]{
		store: make(map[string]cacheItem[T]),
		opts:  opts,
	}
	go c.startAutoRefresh(ctx)
	return c
}
