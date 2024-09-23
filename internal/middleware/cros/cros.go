package cros

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CorsByRules 按照白名单规则设置跨域
func CorsByRules(whiteList []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("origin")
		for _, v := range whiteList {
			if v == origin {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Headers",
					"Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, X-Token, X-User-Id")
				c.Header("Access-Control-Allow-Methods", "POST, GET")
				c.Header("Access-Control-Expose-Headers",
					"Content-Length, Access-Control-Allow-Origin, "+
						"Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
				c.Header("Access-Control-Allow-Credentials", "true")
				break
			}
		}

		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
			return
		}

		c.Next()
	}
}
