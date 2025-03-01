package caches

import (
	"go_admin_api/global"
)

// Cache 定义缓存接口

// CacheType 定义缓存类型
type CacheType string

const (
	CacheTypeBigCache string = "bigcache"
	CacheTypeRedis    string = "redis"
)

type Cacheinit struct{}

// NewCache 根据配置创建缓存实例
func (f *Cacheinit) NewCache(cacheType string, cnf *global.Application) (global.Cache, error) {
	switch cacheType {
	case CacheTypeBigCache:
		return NewBigCache()
	case CacheTypeRedis:
		addr := cnf.CacheConfig.CACHE_HOST
		password := cnf.CacheConfig.CACHE_PASSWPRD
		db := cnf.CacheConfig.CACHE_DB
		return NewRedisCache(addr, password, db)
	default:
		return NewBigCache()
	}
}

func InitNewCache() error {
	cacheInit := &Cacheinit{}

	// 从配置中获取缓存类型
	cacheType := global.App.CacheConfig.CACHE_TYPE

	// 创建缓存实例
	cacheInstance, err := cacheInit.NewCache(cacheType, global.App)
	if err != nil {
		return err
	}

	// 将缓存实例保存到全局变量中
	global.App.Cache = cacheInstance

	return nil
}
