package user

import (
	"fmt"
	"time"

	"github.com/TensoRaws/NuxBT-Backend/internal/common/cache"
	"github.com/TensoRaws/NuxBT-Backend/internal/common/db"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
)

type InvitationGenResponse struct {
	InvitationCode string `json:"invitation_code"`
}

type InvitationMeResponse []cache.UserInvitation

// InvitationGen 生成邀请码 (POST /invitation/gen)
func InvitationGen(c *gin.Context) {
	userID, _ := resp.GetUserIDFromGinContext(c)

	user, err := db.GetUserByID(userID)
	if err != nil {
		resp.AbortWithMsg(c, code.DatabaseErrorRecordNotFound, "User not found!")
		return
	}

	// 检查用户是否有资格生成邀请码
	if time.Now().Unix() < user.CreatedAt.Unix()+int64(config.RegisterConfig.InvitationCodeEligibilityTime*3600*24) {
		resp.AbortWithMsg(c, code.UserErrorInvitationCodeEligibilityTimeNotReached,
			fmt.Sprintf("You need to wait %d days after registration to generate invitation code!",
				config.RegisterConfig.InvitationCodeEligibilityTime))
		log.Logger.Infof("User %d tried to generate invitation code before eligibility time!", userID)
		return
	}

	// 检查用户邀请码数量
	count, err := cache.GetValidInvitationCodeCountByUserID(userID)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		return
	}
	log.Logger.Infof("User %d has %d valid invitation codes!", userID, count)

	// 检查用户邀请码数量是否达到上限
	if count >= config.RegisterConfig.InvitationCodeLimit {
		resp.AbortWithMsg(c, code.UserErrorInvitationCodeHasReachedLimit,
			fmt.Sprintf("You have generated %d invitation codes!", count))
		return
	}

	// 生成邀请码
	codeGen, err := cache.GenerateInvitationCode(userID)
	if err != nil {
		resp.AbortWithMsg(c, code.UnknownError, err.Error())
		return
	}

	resp.OKWithData(c, InvitationGenResponse{InvitationCode: codeGen})
	log.Logger.Infof("User %d generated invitation code_gen %s successfully!", userID, codeGen)
}

// InvitationMe 获取邀请码列表 (GET /invitation/me)
func InvitationMe(c *gin.Context) {
	userID, _ := resp.GetUserIDFromGinContext(c)

	codeList, err := cache.GetInvitationCodeListByUserID(userID)
	if err != nil {
		return
	}

	if len(codeList) == 0 {
		resp.OKWithData(c, InvitationMeResponse{})
	} else {
		resp.OKWithData(c, InvitationMeResponse(codeList))
	}
	log.Logger.Infof("User %d got invitation code list successfully!", userID)
}
