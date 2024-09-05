package settings_api

import (
	"gvb_server/global"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}
// SettingsInfoView 获取配置信息
// @Tags 系统管理
// @Summary 系统配置查看
// @Description 获取指定类型的配置信息
// @ID get-settings-info
// @Param name path string true "配置类型名称"
// @Router /api/settings/{name} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var uri SettingsUri
	// 绑定uri参数
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	name := uri.Name
	//name := c.Param("name")
	switch name {
	case "siteinfo":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "upload":
		res.OkWithData(global.Config.Upload, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
}
