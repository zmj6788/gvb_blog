package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/untils/jwts"
	"gvb_server/untils/pwd"
	"gvb_server/untils/random"

	"github.com/gin-contrib/sessions"
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

	session := sessions.Default(c)
	if cr.Code == nil {
		// 第一次提交，后台发送验证码
		// 生成4位验证码，
		code := random.Code(4)
		// 将验证码保存到session中
		// session 数据是存储在客户端的 cookie 中
		session.Set("email_code", code)
		session.Set("first_email", cr.Email)

		err := session.Save()
		if err != nil {
			global.Log.Error("session 保存失败", err)
			res.FailWithMessage("session 错误", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "你的验证码为:"+code)
		if err != nil {
			global.Log.Error("验证码发送失败", err)
			res.FailWithMessage("验证码发送失败", c)
			return
		}
		res.OkWithMessage("验证码已发送，请注意查收", c)
		return
	}

	//第二次提交邮箱账号以及验证码和密码，用来绑定邮箱
	code := session.Get("email_code")
	email := session.Get("first_email")
	// 判断邮箱与第一次提交的邮箱是否一致	验证码是否一致
	if cr.Email != email {
		res.FailWithMessage("请保证与收取验证码的邮箱一致", c)
		return
	}
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	if len(cr.Password) < 6  {
		res.FailWithMessage("密码安全性过低", c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		global.Log.Error("用户不存在", err)
		res.FailWithMessage("用户不存在", c)
		return
	}
	
	hashPwd := pwd.HashPwd(cr.Password)
	
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	
	if err != nil {
		global.Log.Error("邮箱绑定失败", err)
		res.FailWithMessage("邮箱绑定失败", c)
		return
	}
	res.OkWithMessage("邮箱绑定成功", c)
}
