package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
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
		res.FailWithMessage("账号不存在", c)
		return
	}

	//验证密码是否正确
	isCheck := pwd.CheckPwd(userModel.Password, req.Password)
	if !isCheck {
		global.Log.Warn("密码错误")
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
	//更新数据库token,使得管理员更新用户权限时可以获取到用户token使该token失效
	//达到使用户强制重新登陆，更新用户权限的功能
	global.DB.Model(&userModel).Update("token", token)
	if err != nil {
		global.Log.Error("生成token失败", err.Error())
		res.FailWithMessage("生成token失败", c)
		return
	}
	res.Ok(gin.H{"token":token}, "登录成功", c)
}
