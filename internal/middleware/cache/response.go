package cache

import (
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
	"time"
)

// Response 缓存接口响应的中间件
func Response(redisClient *cache.Client, ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成缓存键，使用请求的 URL 和方法
		cacheKey := c.Request.Method + ":" + c.Request.URL.String()

		// 尝试从缓存中获取响应
		cachedResponse, err := redisClient.Get(cacheKey).Result()
		if err == nil {
			// 缓存命中，直接返回缓存的响应
			util.OKWithCache(c, cachedResponse)
			log.Logger.Debug("Cache hit: " + cacheKey)
			return
		}

		// 缓存未命中，调用后续的处理函数
		c.Next()

		// 调用结束后，将结果存入缓存
		result, exists := c.Get("cache")
		if exists {
			redisClient.Set(cacheKey, result, ttl)
		}
	}
}
