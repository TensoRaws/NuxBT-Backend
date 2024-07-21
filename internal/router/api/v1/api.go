package v1

import (
	"net/http"
	"time"

	middleware_cache "github.com/TensoRaws/NuxBT-Backend/internal/middleware/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/cros"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/logger"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/gin-gonic/gin"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(cros.CorsByRules(config.ServerConfig.Cros))                                        // 跨域中间件
	r.Use(logger.DefaultLogger(), gin.Recovery())                                            // 日志中间件
	r.Use(middleware_cache.NewRateLimiter(config.ServerConfig.RequestLimit, 60*time.Second)) // 限流中间件

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "NuxBT-Backend",
		})
	})

	api := r.Group("/api/v1/")
	{
		UserRoutes(api)
	}

	return r
}
