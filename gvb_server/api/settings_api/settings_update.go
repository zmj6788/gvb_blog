package settings_api

import (
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)
// SettingsInfoUpdateView 更新配置信息
// @Tags 系统管理
// @Summary 更新配置信息
// @Description 更新配置信息
// @ID update-settings-info
// @Param name path string true "配置类型名称"
// @Param data body config.SiteInfo false "站点信息配置"
// @Param data body config.Email false "邮箱配置"
// @Param data body config.Jwt false "JWT 配置"
// @Param data body config.QiNiu false "七牛云配置"
// @Param data body config.QQ false "QQ 配置"
// @Param data body config.Upload false "上传配置"
// @Router /api/settings/{name} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	name := c.Param("name")
	switch name {
	case "siteinfo":
		var si config.SiteInfo
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.SiteInfo = si
	case "email":
		var si config.Email
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.Email = si
	case "jwt":
		var si config.Jwt
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.Jwt = si
	case "qiniu":
		var si config.QiNiu
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.QiNiu = si
	case "qq":
		var si config.QQ
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.QQ = si
	case "upload":
		var si config.Upload
		//将请求体中json数据绑定到si结构体中
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		//配置信息修改
		global.Config.Upload = si
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}

	//配置信息写入到yaml中(储存)
	err := core.SetYaml()
	if err != nil {
		//服务端查看
		logrus.Error(err)
		//返回客户端
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("修改成功", c)
}
