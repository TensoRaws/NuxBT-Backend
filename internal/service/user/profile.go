package user

import (
	"strconv"

	"github.com/TensoRaws/NuxBT-Backend/internal/common/dao"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
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
	Roles      []string `json:"roles,omitempty"`
	Signature  string   `json:"signature"`
	UserID     int32    `json:"user_id"`
	Username   string   `json:"username"`
}

type ProfileOthersRequest struct {
	UserID int32 `form:"user_id" binding:"required"`
}

// ProfileMe 获取用户自己的信息 (GET /profile/me)
func ProfileMe(c *gin.Context) {
	userID, _ := util.GetUserIDFromGinContext(c)

	user, err := dao.GetUserByID(userID)
	if err != nil {
		util.AbortWithMsg(c, "User not found")
		return
	}

	roles, err := dao.GetUserRolesByID(userID)
	if err != nil {
		log.Logger.Info("Failed to get user roles: " + err.Error())
		roles = []string{}
	}

	util.OKWithData(c, ProfileResponse{
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

	log.Logger.Info("get user profile success: " + util.StructToString(user))
}

// ProfileOthers 用户查询他人信息 (GET /profile)
func ProfileOthers(c *gin.Context) {
	// 绑定参数
	var req ProfileOthersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.AbortWithMsg(c, "invalid request: "+err.Error())
		return
	}

	userID, _ := util.GetUserIDFromGinContext(c)

	// 获取信息
	user, err := dao.GetUserByID(req.UserID)
	if err != nil {
		util.AbortWithMsg(c, "User not found")
		return
	}

	roles, err := dao.GetUserRolesByID(req.UserID)
	if err != nil {
		log.Logger.Info("Failed to get user roles: " + err.Error())
		roles = []string{}
	}
	// 判断是否为隐私账号
	if user.Private {
		// 只显示最基础信息
		util.OKWithData(c, ProfileResponse{
			Avatar:     user.Avatar,
			Background: user.Background,
			CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
			Email:      nil,
			Experience: nil,
			Inviter:    nil,
			LastActive: user.LastActive.Format("2006-01-02 15:04:05"),
			Private:    true,
			Roles:      nil,
			Signature:  user.Signature,
			UserID:     user.UserID,
			Username:   user.Username,
		})
	} else {
		// 显示全部信息
		util.OKWithData(c, ProfileResponse{
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

	log.Logger.Info("Get user profile success: " + strconv.Itoa(int(req.UserID)) +
		", by user ID: " + strconv.Itoa(int(userID)))
}
