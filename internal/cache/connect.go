package cache

import (
	redis "github.com/go-redis/redis/v8"
)

type Cache struct {
	Client *redis.Client
}

type CacheConfig struct {
	Addr string
}

func NewLocalConfig(Adrr string) *CacheConfig {
	return &CacheConfig{
		Addr: Adrr,
	}
}

func NewCache(cn *CacheConfig) *Cache {
	return &Cache{
		Client: redis.NewClient(&redis.Options{
			Addr: cn.Addr,
		}),
	}
}
