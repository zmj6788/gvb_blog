package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func TagRouter(router *gin.RouterGroup) {
	TagApi := api.ApiGroupApp.TagApi

	router.POST("/tags", TagApi.TagCreateView)
	router.DELETE("/tags", TagApi.TagRemoveView)
	router.PUT("/tags/:id", TagApi.TagUpdateView)
	router.GET("/tags", TagApi.TagListView)

}
