package cache

import (
	"fmt"
	"time"

	"github.com/TensoRaws/NuxBT-Backend/module/cache"
	"github.com/TensoRaws/NuxBT-Backend/module/config"
	"github.com/TensoRaws/NuxBT-Backend/module/util"
)

type UserInvitationMapValue struct {
	CreatedAt int64 `json:"created_at"`
	UsedBy    int32 `json:"used_by"`
	ExpiresAt int64 `json:"expires_at"`
}

// GenerateInvitationCode 生成邀请码
func GenerateInvitationCode(userID int32) (string, error) {
	c := cache.Clients[cache.InvitationCode]

	expTime := time.Duration(config.RegisterConfig.InvitationCodeExpirationTime) * time.Hour * 24
	code := util.GetRandomString(24)
	// 将生成的邀请码存储到 Redis
	err := c.Set(code, userID, expTime).Err()
	if err != nil {
		return "", err
	}

	toMapString := util.StructToString(UserInvitationMapValue{
		CreatedAt: time.Now().Unix(),              // 存储邀请码的创建时间
		UsedBy:    0,                              // 初始状态为未使用
		ExpiresAt: time.Now().Add(expTime).Unix(), // 过期时间
	})

	// 将邀请码信息存储到用户的哈希表中，方便查询
	err = c.HMSet(fmt.Sprintf("user:%d:invitations", userID), map[string]any{code: toMapString}).Err()
	if err != nil {
		return "", err
	}

	// 更新哈希表键的过期时间，为 10 倍的邀请码过期时间，保证一段时间内可以查询到邀请码状态
	err = c.Expire(fmt.Sprintf("user:%d:invitations", userID), 10*expTime).Err()
	if err != nil {
		return "", err
	}

	return code, nil
}

type UserInvitation struct {
	InvitationCode string `json:"invitation_code"`
	UserInvitationMapValue
}

// GetInvitationCodeByUserID 获取用户近期的邀请码信息
func GetInvitationCodeByUserID(userID int32) ([]UserInvitation, error) {
	c := cache.Clients[cache.InvitationCode]

	// 从 Redis 中获取用户的邀请码信息
	invitations, err := c.HGetAll(fmt.Sprintf("user:%d:invitations", userID)).Result()
	if err != nil {
		return nil, err
	}

	var invitationList []UserInvitation
	for code, info := range invitations {
		var uim UserInvitationMapValue
		err := util.StringToStruct(info, &uim)
		if err != nil {
			return nil, err
		}
		invitationList = append(invitationList, UserInvitation{
			InvitationCode:         code,
			UserInvitationMapValue: uim,
		})
	}

	return invitationList, nil
}

// GetValidInvitationCodeCountByUserID 获取用户有效的邀请码数量
func GetValidInvitationCodeCountByUserID(userID int32) (int, error) {
	c := cache.Clients[cache.InvitationCode]

	invitations, err := c.HGetAll(fmt.Sprintf("user:%d:invitations", userID)).Result()
	if err != nil {
		return 0, err
	}

	count := 0
	for _, info := range invitations {
		var uim UserInvitationMapValue
		err := util.StringToStruct(info, &uim)
		if err != nil {
			return 0, err
		}

		if uim.UsedBy == 0 && uim.ExpiresAt > time.Now().Unix() {
			count++
		}
	}

	return count, nil
}
