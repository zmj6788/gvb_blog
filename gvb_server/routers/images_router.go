package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func ImagesRouter(router *gin.RouterGroup) {
	ImagesApi := api.ApiGroupApp.ImagesApi
	router.POST("/images", ImagesApi.ImageUploadView)
}
