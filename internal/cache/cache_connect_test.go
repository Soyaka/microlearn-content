package cache

import "testing"

func TestCacheConnection(t *testing.T) {
	testConfig := NewCacheConfig("localhost:6379")
	cache := NewCache(testConfig)
	if cache == nil || cache.Client == nil {
		t.Errorf("Cache connection failed")
	}
	defer cache.Client.Close()
}
