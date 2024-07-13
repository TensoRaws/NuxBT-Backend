package cache

import (
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"time"

	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/third_party/gin_cache"
	"github.com/TensoRaws/NuxBT-Backend/third_party/gin_cache/persist"
	"github.com/gin-gonic/gin"
)

// Response 缓存接口响应的中间件
func Response(redisClient *cache.Client, ttl time.Duration, queryFilter []string) gin.HandlerFunc {
	redisStore := persist.NewRedisStore(redisClient.C)

	var strategy gin_cache.Option
	if queryFilter != nil {
		strategy = gin_cache.WithCacheStrategyByRequest(func(c *gin.Context) (bool, gin_cache.Strategy) {
			// 剔除 query 参数
			return true, gin_cache.Strategy{
				CacheKey: util.RemoveQueryParameter(c.Request.RequestURI, queryFilter...),
			}
		})
	} else {
		strategy = gin_cache.WithCacheStrategyByRequest(func(c *gin.Context) (bool, gin_cache.Strategy) {
			return true, gin_cache.Strategy{
				CacheKey: c.Request.RequestURI,
			}
		})
	}

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
