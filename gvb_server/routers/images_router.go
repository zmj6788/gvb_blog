package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func ImagesRouter(router *gin.RouterGroup) {
	ImagesApi := api.ApiGroupApp.ImagesApi
	router.GET("/images", ImagesApi.ImageListView)
	router.POST("/images", ImagesApi.ImageUploadView)
	router.DELETE("/images", ImagesApi.ImageRemoveView)
}
