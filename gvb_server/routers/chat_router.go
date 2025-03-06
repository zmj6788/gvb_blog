package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func ChatRouter(router *gin.RouterGroup) {
	ChatApi := api.ApiGroupApp.ChatApi
	router.GET("/chat_groups", ChatApi.ChatGroupView)

}
