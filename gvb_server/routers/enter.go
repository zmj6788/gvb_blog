package routers

import (
	"gvb_server/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
<<<<<<< HEAD
	//如有需求在这里读取json错误码文件

	routerGroup := router.Group("/api")
	SettingsRouter(routerGroup)
	ImagesRouter(routerGroup)
=======
	//在这里读取json错误码文件

	routerGroup := router.Group("/api")
	SettingsRouter(routerGroup)
>>>>>>> 2f9e4d1a6a0ab0002a002517dace0301441cd6ca
	return router
}
