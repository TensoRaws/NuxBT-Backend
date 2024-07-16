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
	Expiration int64  `json:"expiration"`
	Token      string `json:"token"`
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
		token := jwt.GenerateToken(user.UserID)

		claims, err := jwt.ParseToken(token)
		if err != nil {
			resp.AbortWithMsg(c, code.UnknownError, err.Error())
			return
		}

		resp.OKWithData(c, LoginResponse{
			Expiration: claims.ExpiresAt.Unix(),
			Token:      token,
		})
	} else {
		resp.Abort(c, code.UserErrorInvalidPassword)
		return
	}
}

// TokenRefresh 用户刷新 token (POST /token/refresh)
func TokenRefresh(c *gin.Context) {
	userID, _ := resp.GetUserIDFromGinContext(c)

	token := jwt.GenerateToken(userID)

	claims, err := jwt.ParseToken(token)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		return
	}

	resp.OKWithData(c, LoginResponse{
		Expiration: claims.ExpiresAt.Unix(),
		Token:      token,
	})
}
