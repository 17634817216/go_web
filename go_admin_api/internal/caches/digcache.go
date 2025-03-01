package caches

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"time"
)

// BigCache 实现 Cache 接口
type BigCache struct {
	client *bigcache.BigCache
}

// NewBigCache 创建 BigCache 实例
func NewBigCache() (*BigCache, error) {
	client, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		return nil, err
	}
	return &BigCache{client: client}, nil
}

func (b *BigCache) Get(ctx context.Context, key string) ([]byte, error) {
	return b.client.Get(key)
}

func (b *BigCache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return b.client.Set(key, value)
}

func (b *BigCache) Delete(ctx context.Context, key string) error {
	return b.client.Delete(key)
}

func (b *BigCache) Close() error {
	return b.client.Close()
}
