package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func AdvertRouter(router *gin.RouterGroup) {
	AdvertApi := api.ApiGroupApp.AdvertApi

	router.POST("/adverts", AdvertApi.AdvertCreateView)
	router.DELETE("/adverts", AdvertApi.AdvertRemoveView)
	router.PUT("/adverts/:id", AdvertApi.AdvertUpdateView)
	router.GET("/adverts", AdvertApi.AdvertListView)

}
