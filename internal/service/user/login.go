package user

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/dao"
	"github.com/TensoRaws/NuxBT-Backend/internal/middleware/jwt"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Login 用户登录 (POST /login)
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	// GORM 查询
	user, err := dao.GetUserByEmail(req.Email)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, "User not found")
		return
	}

	// verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err == nil {
		// 注册之后的下次登录成功，才会为其生成 token
		token := jwt.GenerateToken(user)
		// 打印相应信息和用户信息以及生成的 token 值
		resp.OKWithData(c, LoginResponse{
			Token: token,
		})
	} else {
		resp.Abort(c, code.UserErrorInvalidPassword)
		return
	}
}
