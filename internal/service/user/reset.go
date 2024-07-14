package user

import (
	"fmt"

	"github.com/TensoRaws/NuxBT-Backend/module/util"
	"github.com/gin-gonic/gin"
)

type SetRequest struct {
	Newpassword string `json:"new_password" binding:"required"`
}

// 重置密码 (POST /password/reset)
func ReSetPass(c *gin.Context) {
	//绑定参数
	var req SetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.AbortWithMsg(c, "invalid request")
		return
	}
	fmt.Println(req.Newpassword)
	//鉴权
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

	//修改密码
	err = SetUserPass(user, req.Newpassword)
	if err != nil {
		util.AbortWithMsg(c, "reset password fail")
	}
	//返回
	util.OKWithMsg(c, "reset password success")

}
