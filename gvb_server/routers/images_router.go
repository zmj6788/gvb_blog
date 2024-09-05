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
	router.PUT("/images", ImagesApi.ImageUpdateView)

	//静态代理图片资源
	//使得可以通过如下方式访问图片资源，前提路径完整
	//http://localhost:8080/api/uploads/file/1.jpg
	router.Static("/uploads/file", "./uploads/file")
}
