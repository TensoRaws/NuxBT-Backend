package user

import (
	"strconv"

	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

type ProfileMeResponse struct {
	Avatar     string   `json:"avatar"`
	Background string   `json:"background"`
	CreatedAt  string   `json:"created_at"`
	Email      string   `json:"email"`
	Experience string   `json:"experience"`
	Inviter    string   `json:"inviter"`
	LastActive string   `json:"last_active"`
	Private    bool     `json:"private"`
	Roles      []string `json:"roles"`
	Signature  string   `json:"signature"`
	UserID     string   `json:"user_id"`
	Username   string   `json:"username"`
}

// ProfileMe 获取用户自己的信息 (GET /profile/me)
func ProfileMe(c *gin.Context) {
	userID, err := util.GetUserIDFromGinContext(c)
	if err != nil {
		util.AbortWithMsg(c, "Please login first")
		return
	}

	user, err := GetUserByID(int32(userID))
	if err != nil {
		util.AbortWithMsg(c, "User not found")
		return
	}

	roles, err := GetUserRolesByID(int32(userID))
	if err != nil {
		log.Logger.Info("Failed to get user roles: " + err.Error())
		roles = []string{}
	}

	util.OKWithDataStruct(c, ProfileMeResponse{
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
		UserID:     strconv.Itoa(int(user.UserID)),
		Username:   user.Username,
	})

	log.Logger.Info("get user profile success: " + util.StructToString(user))
}
