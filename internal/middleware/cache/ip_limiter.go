package cache

import (
	"time"

	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	redisLimiter "github.com/ulule/limiter/v3/drivers/store/redis"
)

func NewRateLimiter(redisClient *cache.Client, limit int, slidingWindow time.Duration) gin.HandlerFunc {
	rate := limiter.Rate{
		Period: slidingWindow,
		Limit:  int64(limit),
	}
	store, err := redisLimiter.NewStore(redisClient.C)
	if err != nil {
		log.Logger.Error(err)
	}
	l := limiter.New(store, rate, limiter.WithClientIPHeader("True-Client-IP"))

	return mgin.NewMiddleware(l)
}
