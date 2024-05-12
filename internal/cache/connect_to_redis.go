package cache

import (
	"fmt"
	"log"

	redis "github.com/go-redis/redis/v8"
)

type Cache struct {
	Client *redis.Client
}

type CacheConfig struct {
	Addr string
}

func NewCacheConfig(Adrr string) *CacheConfig {
	return &CacheConfig{
		Addr: Adrr,
	}
}

func NewCache(cn *CacheConfig) *Cache {
	Cache := &Cache{
		Client: redis.NewClient(&redis.Options{
			Addr: cn.Addr,
		}),
	}
	if Cache.Client == nil {
		log.Fatalf("Cache connection failed")
		return nil
	}
	fmt.Println("Connected to Redis!", Cache.Client)
	return Cache
}
