package routers

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//如有需求在这里读取json错误码文件

	routerGroup := router.Group("/api")
	SettingsRouter(routerGroup)
	ImagesRouter(routerGroup)
	AdvertRouter(routerGroup)
	return router
}
