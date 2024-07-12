package v1

import (
	"net/http"
	"time"

	middleware_cache "github.com/TensoRaws/NuxBT-Backend/internal/middleware/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/logger"
	user_service "github.com/TensoRaws/NuxBT-Backend/internal/service/user"
	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/gin-gonic/gin"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(logger.DefaultLogger(), gin.Recovery()) // 日志中间件
	r.Use(middleware_cache.NewRateLimiter(
		cache.Clients[cache.IPLimit], config.ServerConfig.RequestLimit, 60*time.Second))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "NuxBT-Backend",
		})
	})

	api := r.Group("/api/v1/")
	{
		user := api.Group("user/")
		{
			// 用户注册
			user.POST("register", user_service.Register)
			// 用户登录
			user.POST("login", user_service.Login)
			// 用户信息
			user.GET("profile/me", jwt.RequireAuth(), user_service.ProfileMe)
		}
	}

	return r
}
