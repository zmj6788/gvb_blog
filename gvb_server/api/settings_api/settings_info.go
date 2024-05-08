package settings_api

import (
<<<<<<< HEAD
	"gvb_server/global"
=======
>>>>>>> 2f9e4d1a6a0ab0002a002517dace0301441cd6ca
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

<<<<<<< HEAD
type SettingsUri struct {
	Name string `uri:"name"`
}
// SettingsInfoView 获取配置信息
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
=======
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(1001, c)
>>>>>>> 2f9e4d1a6a0ab0002a002517dace0301441cd6ca
}
