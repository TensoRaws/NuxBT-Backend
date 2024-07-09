package v1

import (
	middleware_cache "github.com/TensoRaws/NuxBT-Backend/internal/middleware/cache"
	"net/http"
	"time"

	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/gin-gonic/gin"

	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/logger"
	user_service "github.com/TensoRaws/NuxBT-Backend/internal/service/user"
)

func NewAPI() *gin.Engine {
	r := gin.New()
	r.Use(logger.DefaultLogger(), gin.Recovery()) // 日志中间件
	r.Use(middleware_cache.NewRateLimiter(
		middleware_cache.Clients[cache.IPLimit], "general", 200, 60*time.Second))

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
			user.POST("register/", user_service.Register)
			// 用户登录
			user.POST("login/", user_service.Login)
		}
	}

	return r
}
