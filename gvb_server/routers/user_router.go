package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	userApis := api.ApiGroupApp.UserApi
	// 邮箱登录
	router.POST("/email_login", userApis.EmailLoginView)
	//只有登录的用户才能调用用户信息列表
	router.GET("/users", middleware.JwtAuth(), userApis.UserListView)
	router.PUT("/user_pwd",middleware.JwtAuth(),userApis.UserUpdatePasswordView)
	router.POST("/logout",userApis.UserLogoutView)
	//管理员才能修改用户权限
	router.PUT("/user_role",middleware.JwtAdmin(),userApis.UserUpdateRoleView)
	router.DELETE("/users",middleware.JwtAdmin(),userApis.UserRemoveView)
}
