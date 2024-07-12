package user

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Login 用户登录 (POST /login)
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.AbortWithMsg(c, "invalid request")
		return
	}

	// GORM 查询
	user, err := GetUserByEmail(req.Email)
	if err != nil {
		util.AbortWithMsg(c, "User not found")
		return
	}

	// verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err == nil {
		// 注册之后的下次登录成功，才会为其生成 token
		token := jwt.GenerateToken(user)
		// 打印相应信息和用户信息以及生成的 token 值
		util.OKWithData(c, false, LoginResponse{
			Token: token,
		})
	} else {
		util.AbortWithMsg(c, "Invalid Username or Password")
		return
	}
}
