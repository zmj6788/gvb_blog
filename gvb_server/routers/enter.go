package routers

import (
	"gvb_server/global"
  gs "github.com/swaggo/gin-swagger"
	 swaggerFiles "github.com/swaggo/files"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	// swagger使用
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//如有需求在这里读取json错误码文件

	routerGroup := router.Group("/api")
	SettingsRouter(routerGroup)
	ImagesRouter(routerGroup)
	AdvertRouter(routerGroup)
	MenuRouter(routerGroup)
	UserRouter(routerGroup)
	return router
}
