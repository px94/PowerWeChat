package kernel

import (
	"time"

	"github.com/ArtisanCloud/PowerLibs/v3/cache"
	"github.com/redis/go-redis/v9"
)

type InteractsWithCache struct {
	Cache CacheInterface
}

type CacheInterface cache.CacheInterface
type UniversalOptions redis.UniversalOptions

func NewRedisClient(options *UniversalOptions) CacheInterface {
	if options == nil {
		return nil
	}

	if len(options.Addrs) == 0 {
		return nil
	}

	return cache.NewGRedis((*redis.UniversalOptions)(options))
}

func NewInteractsWithCache(client CacheInterface) *InteractsWithCache {

	interact := &InteractsWithCache{
		Cache: client,
	}
	if client == nil {
		interact.Cache = interact.createDefaultCache()
	}

	return interact
}

func (interactCache *InteractsWithCache) GetCache() CacheInterface {
	if interactCache.Cache != nil {
		return interactCache.Cache
	}

	// create default cache
	interactCache.Cache = interactCache.createDefaultCache()

	return interactCache.Cache
}

func (interactCache *InteractsWithCache) SetCache(cache CacheInterface) *InteractsWithCache {

	interactCache.Cache = cache

	return interactCache
}

func (interactCache *InteractsWithCache) createDefaultCache() CacheInterface {

	interactCache.Cache = cache.NewMemCache("ac.go.power", time.Duration(1500), "")

	return interactCache.Cache
}
