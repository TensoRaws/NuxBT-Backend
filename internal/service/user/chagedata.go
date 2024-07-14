package user

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/dao"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

type ChangeUserRequest struct {
	Avatar     string `json:"avatar" `
	Background string `json:"background" `
	Email      string `json:"email" `
	Private    int    `json:"private" `
	Signature  string `json:"signature" `
	Username   string `json:"username"`
}

// ChangeUser 用户信息更新 (POST /profile/update)
func ChangeUser(c *gin.Context) {
	// 参数绑定
	var req ChangeUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.AbortWithMsg(c, "invalid request")
		return
	}

	userID, err := util.GetUserIDFromGinContext(c)
	if err != nil {
		util.AbortWithMsg(c, "Please login first")
		return
	}

	user, err := dao.GetUserByID(int32(userID))
	if err != nil {
		util.AbortWithMsg(c, "User not found")
		return
	}
	// 迁移数据
	updates := make(map[string]interface{})
	if req.Private == 1 {
		updates["private"] = true
	} else if req.Private == 2 {
		updates["private"] = false
	} else if req.Private != 0 {
		util.AbortWithMsg(c, "invalid request")
		return
	}
	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.Signature != "" {
		updates["signature"] = req.Signature
	}
	if req.Background != "" {
		updates["background"] = req.Background
	}
	// 执行更新
	dao.UpdateUserData(user, updates)
	util.OKWithMsg(c, "update success")
}
