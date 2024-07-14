package user

import (
	"strconv"

	"github.com/TensoRaws/NuxBT-Backend/internal/service/common/dao"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

type ProfileResponse struct {
	Avatar     string   `json:"avatar"`
	Background string   `json:"background"`
	CreatedAt  string   `json:"created_at"`
	Email      string   `json:"email"`
	Experience string   `json:"experience"`
	Inviter    string   `json:"inviter"`
	LastActive string   `json:"last_active"`
	Private    bool     `json:"private"`
	Roles      []string `json:"roles,omitempty"`
	Signature  string   `json:"signature"`
	UserID     int      `json:"user_id"`
	Username   string   `json:"username"`
}

type UserProfileRequest struct {
	UserId int `form:"user_id" binding:"required"`
}

// ProfileMe 获取用户自己的信息 (GET /profile/me)
func ProfileMe(c *gin.Context) {
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

	roles, err := dao.GetUserRolesByID(int32(userID))
	if err != nil {
		log.Logger.Info("Failed to get user roles: " + err.Error())
		roles = []string{}
	}

	util.OKWithData(c, ProfileResponse{
		Avatar:     user.Avatar,
		Background: user.Background,
		CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
		Email:      user.Email,
		Experience: strconv.Itoa(int(user.Experience)),
		Inviter:    strconv.Itoa(int(user.Inviter)),
		LastActive: user.LastActive.Format("2006-01-02 15:04:05"),
		Private:    user.Private,
		Roles:      roles,
		Signature:  user.Signature,
		UserID:     int(user.UserID),
		Username:   user.Username,
	})

	log.Logger.Info("get user profile success: " + util.StructToString(user))
}

// ProfileOthers 用户查询他人信息 (GET /profile)
func ProfileOthers(c *gin.Context) {
	// 绑定参数
	var req UserProfileRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.AbortWithMsg(c, "invalid request")
		return
	}
	// 鉴权
	userID, err := util.GetUserIDFromGinContext(c)
	if err != nil {
		util.AbortWithMsg(c, "Please login first")
		return
	}
	// 仅用于鉴权不使用
	_, err = dao.GetUserByID(int32(userID))
	if err != nil {
		util.AbortWithMsg(c, "User not found")
		return
	}
	// 获取信息
	user, err := dao.GetUserByID(int32(req.UserId))
	if err != nil {
		util.AbortWithMsg(c, "User not found")
		return
	}

	roles, err := dao.GetUserRolesByID(int32(userID))
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
			CreatedAt:  "",
			Email:      "",
			Experience: "",
			Inviter:    "",
			LastActive: "",
			Private:    user.Private,
			Roles:      []string{},
			Signature:  "",
			UserID:     int(user.UserID),
			Username:   user.Username,
		})
	} else {
		// 显示全部信息

		util.OKWithData(c, ProfileResponse{
			Avatar:     user.Avatar,
			Background: user.Background,
			CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
			Email:      user.Email,
			Experience: strconv.Itoa(int(user.Experience)),
			Inviter:    strconv.Itoa(int(user.Inviter)),
			LastActive: user.LastActive.Format("2006-01-02 15:04:05"),
			Private:    user.Private,
			Roles:      roles,
			Signature:  user.Signature,
			UserID:     int(user.UserID),
			Username:   user.Username,
		})
	}
}
