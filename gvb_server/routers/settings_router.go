package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func SettingsRouter(router *gin.RouterGroup) {
	settingsApi := api.ApiGroupApp.SettingsApi
<<<<<<< HEAD
	//设置动态路由，便于后期扩展一个接口获取多种配置信息以及更改多种配置信息
	router.GET("/settings/:name", settingsApi.SettingsInfoView)
	router.PUT("/settings/:name", settingsApi.SettingsInfoUpdateView)
=======
	router.GET("/settingsinfo", settingsApi.SettingsInfoView)
>>>>>>> 2f9e4d1a6a0ab0002a002517dace0301441cd6ca
}
