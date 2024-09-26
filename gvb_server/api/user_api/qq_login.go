package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/qq"
	"gvb_server/untils/jwts"
	"gvb_server/untils/pwd"
	"gvb_server/untils/random"

	"github.com/gin-gonic/gin"
)

func (UserApi) QQLoginView(c *gin.Context) {
	
	//code从哪里得到

	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	OpenID := qqInfo.OpenID
	//根据openid判断用户是否存在
	var user models.UserModel
	err = global.DB.Take(&user, "token = ?", OpenID).Error
	if err != nil {
		// 用户不存在
		// 注册
		hashPwd := pwd.HashPwd(random.RandString(16))
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   OpenID, //由于用户可以绑定邮箱登录,这里的用户名随意一点也可以
			Password:   hashPwd,
			Avatar:     qqInfo.Avatar,
			Addr:       "内网",
			Token:      qqInfo.OpenID,
			IP:         c.ClientIP(),
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			res.FailWithMessage("注册失败", c)
			return
		}
	}
	//注册成功自动登录
	//账号存在自动登录
	// 登录操作
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", c)
		return
	}
	res.Ok(gin.H{"token": token}, "登录成功", c)
}
