package user

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// 优先从 url 中获取参数
	username := c.Query("username")
	password := c.Query("password")
}
