package cache

import (
	"time"

	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/TensoRaws/NuxBT-Backend/third_party/gin_cache"
	"github.com/TensoRaws/NuxBT-Backend/third_party/gin_cache/persist"
	"github.com/gin-gonic/gin"
)

// Response 缓存接口响应的中间件，queryFilter 为需要去除的 query 参数，使用其他的来构建缓存 key
func Response(redisClient *cache.Client, ttl time.Duration, queryFilter ...string) gin.HandlerFunc {
	redisStore := persist.NewRedisStore(redisClient.C)

	strategy := gin_cache.WithCacheStrategyByRequest(
		func(c *gin.Context) (bool, gin_cache.Strategy) {
			// 去除 query 参数
			var key string
			if queryFilter != nil {
				key = util.RemoveQueryParameter(c.Request.RequestURI, queryFilter...)
			} else {
				key = c.Request.RequestURI
			}
			return true, gin_cache.Strategy{
				CacheKey: key,
			}
		},
	)

	return gin_cache.CacheByRequestURI(
		redisStore,
		ttl,
		gin_cache.WithOnHitCache(
			func(c *gin.Context) {
				log.Logger.Info("Cache hit: " + c.Request.RequestURI)
			},
		),
		gin_cache.WithOnMissCache(
			func(c *gin.Context) {
				log.Logger.Info("Cache miss, try to cache: " + c.Request.RequestURI)
			},
		),
		strategy,
	)
}
