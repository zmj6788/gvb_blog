package user_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"
	"gvb_server/untils"
	"gvb_server/untils/jwts"
	"gvb_server/untils/pwd"

	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

// EmailLoginView 邮箱登录
// @Tags 用户管理
// @Summary 邮箱登录
// @Description 邮箱登录
// @Param data body EmailLoginRequest    true  "表示多个参数"
// @Router /api/email_login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) EmailLoginView(c *gin.Context) {
	// 日志实例化
	log := log_stash.NewLogByGin(c)
	// 获取邮箱登录参数
	var req EmailLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithError(err, &req, c)
		return
	}

	//验证账号是否存在
	var userModel models.UserModel
	if err := global.DB.Take(&userModel, "user_name = ?  or  email = ?", req.UserName, req.UserName).Error; err != nil {
		global.Log.Warn("账号不存在")
		log.Warn(fmt.Sprintf("%s 账户不存在", req.UserName))
		res.FailWithMessage("账号不存在", c)
		return
	}

	//验证密码是否正确
	isCheck := pwd.CheckPwd(userModel.Password, req.Password)
	if !isCheck {
		global.Log.Warn("密码错误")
		log.Warn(fmt.Sprintf("用户名密码错误 %s , %s", req.UserName, req.Password))
		res.FailWithMessage("密码错误", c)
		return
	}

	//生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   userModel.ID,
		Username: userModel.UserName,
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
	})
	// //更新数据库token,使得管理员更新用户权限时可以获取到用户token使该token失效
	// //达到使用户强制重新登陆，更新用户权限的功能
	// global.DB.Model(&userModel).Update("token", token)
	if err != nil {
		global.Log.Error("生成token失败", err.Error())
		log.Error(fmt.Sprintf("生成token失败 %s ", err.Error()))
		res.FailWithMessage("生成token失败", c)
		return
	}

	ip, addr := untils.GetAddrByGin(c)
	log = log_stash.New(ip, token)
	log.Info("用户登录成功")
	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    c.Request.UserAgent(),
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})

	res.Ok(gin.H{"token": token}, "登录成功", c)
}
