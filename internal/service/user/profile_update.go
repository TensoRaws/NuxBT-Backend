package user

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/dao"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

type ProfileUpdateRequest struct {
	Avatar     *string `json:"avatar" binding:"omitempty"`
	Background *string `json:"background" binding:"omitempty"`
	Email      *string `json:"email" binding:"omitempty,email"`
	Private    *bool   `json:"private" binding:"omitempty"`
	Signature  *string `json:"signature" binding:"omitempty"`
	Username   *string `json:"username" binding:"omitempty"`
}

// ProfileUpdate 用户信息更新 (POST /profile/update)
func ProfileUpdate(c *gin.Context) {
	// 参数绑定
	var req ProfileUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	userID, _ := resp.GetUserIDFromGinContext(c)

	// 准备更新数据
	updates := make(map[string]interface{})

	if req.Private != nil {
		updates["private"] = req.Private
	}

	if req.Username != nil && *req.Username != "" {
		err := util.CheckUsername(*req.Username)
		if err != nil {
			resp.AbortWithMsg(c, code.UserErrorInvalidUsername, err.Error())
			return
		}
		updates["username"] = *req.Username
	}

	if req.Email != nil {
		updates["email"] = *req.Email
	}

	if req.Avatar != nil {
		updates["avatar"] = *req.Avatar
	}

	if req.Signature != nil {
		updates["signature"] = *req.Signature
	}

	if req.Background != nil {
		updates["background"] = *req.Background
	}
	// 执行更新
	err := dao.UpdateUserDataByUserID(userID, updates)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordUpdateFailed, err.Error())
		return
	}

	resp.OK(c)

	log.Logger.Infof("update user profile success, userID: %v", userID)
}
