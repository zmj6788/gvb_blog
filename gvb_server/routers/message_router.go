package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"

	"github.com/gin-gonic/gin"
)

func MessageRouter(router *gin.RouterGroup) {
	MessageApi := api.ApiGroupApp.MessageApi

	router.POST("/messages",middleware.JwtAuth(), MessageApi.MessageCreateView)
	// router.DELETE("/tags", TagApi.TagRemoveView)
	router.PUT("/messages",middleware.JwtAuth() ,MessageApi.MessageRecordView)
	//管理员查看所有消息列表
	router.GET("/messages_all",  middleware.JwtAdmin(),MessageApi.MessageListAllView)
	// 用户查看自己的消息列表,显示与对方最新的一条消息
	router.GET("/messages",  middleware.JwtAuth(),MessageApi.MessageListView)
	

}
