package user

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	username := c.Query("email")
	password := c.Query("password")

	if len(username) == 0 && len(password) == 0 {
		username = c.Request.PostFormValue("email")
		password = c.Request.PostFormValue("password")
	}

	if len(username) == 0 || len(password) == 0 {
		util.AbortWithMsg(c, "username and password is required...")
		return
	}

	// GORM 查询

	// verify password
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	err := bcrypt.CompareHashAndPassword([]byte("fdasa"), []byte(password))
	if err == nil {
		// 注册之后的下次登录成功，才会为其生成 token
		token := jwt.GenerateToken(username)
		// 打印相应信息和用户信息以及生成的 token 值
		util.OKWithData(c, map[string]interface{}{
			"user_id": 1,
			"token":   token,
		})
	} else {
		util.AbortWithMsg(c, "Invalid Username or Password")
		return
	}
}
