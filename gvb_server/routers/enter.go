package routers

import (
	"gvb_server/api/user_api"
	"gvb_server/global"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	// swagger使用
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 测试qq登录接口
	router.GET("/login", user_api.UserApi{}.QQLoginView)
	//如有需求在这里读取json错误码文件

	routerGroup := router.Group("/api")
	SettingsRouter(routerGroup)
	ImagesRouter(routerGroup)
	AdvertRouter(routerGroup)
	MenuRouter(routerGroup)
	UserRouter(routerGroup)
	TagRouter(routerGroup)
	MessageRouter(routerGroup)
	return router
}
