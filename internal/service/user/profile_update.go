package user

import (
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

type ProfileUpdateRequest struct {
	Avatar     string `json:"avatar" binding:"required"`
	Background string `json:"background" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Private    *bool  `json:"private" binding:"required"`
	Signature  string `json:"signature" binding:"required"`
	Username   string `json:"username" binding:"required"`
}

// ProfileUpdate 用户信息更新 (POST /profile/update)
func ProfileUpdate(c *gin.Context) {
	// 参数绑定
	var req ProfileUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	err := util.CheckUsername(req.Username)
	if err != nil {
		resp.AbortWithMsg(c, code.UserErrorInvalidUsername, err.Error())
		return
	}

	userID, _ := resp.GetUserIDFromGinContext(c)

	// 没传数字，直接序列化嗯造
	updateInfo, err := util.StructToMap(req)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		return
	}
	// 执行更新
	err = db.PatchUser(userID, updateInfo)

	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordPatchFailed, err.Error())
		return
	}

	resp.OK(c)

	log.Logger.Infof("update user profile success, userID: %v", userID)
}
