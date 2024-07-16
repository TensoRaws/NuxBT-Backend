package user

import (
	"fmt"

	"github.com/TensoRaws/NuxBT-Backend/internal/common/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/code"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/log"
	"github.com/TensoRaws/NuxBT-Backend/module/resp"
	"github.com/gin-gonic/gin"
)

type InvitationGenResponse struct {
	InvitationCode string `json:"invitation_code"`
}

// InvitationGen 生成邀请码 (POST /invitation/gen)
func InvitationGen(c *gin.Context) {
	userID, _ := resp.GetUserIDFromGinContext(c)

	count, err := cache.GetValidInvitationCodeCountByUserID(userID)
	if err != nil {
		return
	}
	log.Logger.Infof("User %d has %d valid invitation codes!", userID, count)

	if count >= config.RegisterConfig.InvitationCodeLimit {
		resp.AbortWithMsg(c, code.UserErrorInvitationCodeHasReachedLimit,
			fmt.Sprintf("You have generated %d invitation codes!", count))
		return
	}

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

	resp.OKWithData(c, codeList)
	log.Logger.Infof("User %d got invitation code list successfully!", userID)
}
