package user_api

import (
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/untils/jwts"
	"gvb_server/untils/random"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱格式不正确"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	// 用于更新用户邮箱信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	logrus.Info(claims.UserID)

	//第一次提交邮箱账号，用来接收验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	if cr.Code == nil {
		// 第一次提交，后台发送验证码
		// 生成4位验证码，将验证码保存到session中
		
		code := random.Code(4)
		email.NewCode().Send(cr.Email, "你的验证码为:"+code)
		res.OkWithMessage("验证码已发送，请注意查收", c)
		return
	}

	//第二次提交邮箱账号以及验证码和密码，用来绑定邮箱

}
