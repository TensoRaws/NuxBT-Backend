package user

import (
	"time"

	"github.com/TensoRaws/NuxBT-Backend/dal/model"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
)

type ProfileResponse struct {
	Avatar     string   `json:"avatar"`
	Background string   `json:"background"`
	CreatedAt  string   `json:"created_at"`
	Email      *string  `json:"email,omitempty"`
	Experience *int32   `json:"experience,omitempty"`
	Inviter    *int32   `json:"inviter,omitempty"`
	LastActive string   `json:"last_active"`
	Private    bool     `json:"private"`
	Roles      []string `json:"roles"`
	Signature  string   `json:"signature"`
	UserID     int32    `json:"user_id"`
	Username   string   `json:"username"`
}

type ProfileOthersRequest struct {
	UserID int32 `form:"user_id" binding:"required"`
}

// ProfileMe 获取用户自己的信息 (GET /profile/me)
func ProfileMe(c *gin.Context) {
	userID, _ := resp.GetUserIDFromGinContext(c)

	// 更新活跃时间
	err := db.PatchUser(userID, &model.User{LastActive: time.Now()})
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		log.Logger.Error("Failed to update user last active time: " + err.Error())
		return
	}

	user, err := db.GetUserByID(userID)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, "User not found")
		log.Logger.Error("Failed to get user by ID: " + err.Error())
		return
	}

	roles, err := cache.GetUserRolesByID(userID)
	if err != nil {
		log.Logger.Info("Failed to get user roles: " + err.Error())
		roles = []string{}
	}

	resp.OKWithData(c, &ProfileResponse{
		Avatar:     user.Avatar,
		Background: user.Background,
		CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
		Email:      &user.Email,
		Experience: &user.Experience,
		Inviter:    &user.Inviter,
		LastActive: user.LastActive.Format("2006-01-02 15:04:05"),
		Private:    user.Private,
		Roles:      roles,
		Signature:  user.Signature,
		UserID:     user.UserID,
		Username:   user.Username,
	})

	log.Logger.Infof("get user profile success, userID: %v", userID)
}

// ProfileOthers 用户查询他人信息 (GET /profile)
func ProfileOthers(c *gin.Context) {
	// 绑定参数
	var req ProfileOthersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.AbortWithMsg(c, code.RequestErrorInvalidParams, err.Error())
		return
	}

	userID, _ := resp.GetUserIDFromGinContext(c)

	// 获取信息
	user, err := db.GetUserByID(req.UserID)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, "User not found")
		return
	}

	roles, err := cache.GetUserRolesByID(req.UserID)
	if err != nil {
		log.Logger.Info("Failed to get user roles: " + err.Error())
		roles = []string{}
	}
	// 判断是否为隐私账号
	if user.Private {
		// 只显示最基础信息
		resp.OKWithData(c, &ProfileResponse{
			Avatar:     user.Avatar,
			Background: user.Background,
			CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
			Email:      nil,
			Experience: nil,
			Inviter:    nil,
			LastActive: user.LastActive.Format("2006-01-02 15:04:05"),
			Private:    true,
			Roles:      roles,
			Signature:  user.Signature,
			UserID:     user.UserID,
			Username:   user.Username,
		})
	} else {
		// 显示全部信息
		resp.OKWithData(c, &ProfileResponse{
			Avatar:     user.Avatar,
			Background: user.Background,
			CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
			Email:      &user.Email,
			Experience: &user.Experience,
			Inviter:    &user.Inviter,
			LastActive: user.LastActive.Format("2006-01-02 15:04:05"),
			Private:    false,
			Roles:      roles,
			Signature:  user.Signature,
			UserID:     user.UserID,
			Username:   user.Username,
		})
	}

	log.Logger.Infof("Get user %v profile success, by user ID: %v", req.UserID, userID)
}
