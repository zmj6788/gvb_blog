package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	userApis := api.ApiGroupApp.UserApi
	router.POST("/email_login", userApis.EmailLoginView)
	router.GET("/users", userApis.UserListView)
}
