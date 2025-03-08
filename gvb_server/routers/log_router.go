package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func LogRouter(router *gin.RouterGroup) {
	LogApi := api.ApiGroupApp.LogApi

	router.GET("/logs", LogApi.LogListView)
	router.DELETE("/logs", LogApi.LogRemoveListView)

}
