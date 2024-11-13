package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func DiggRouter(router *gin.RouterGroup) {
	DiggApi := api.ApiGroupApp.DiggApi

	router.POST("/digg/article", DiggApi.DiggArticleView)

}
