package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"

	"github.com/gin-gonic/gin"
)

func AdvertRouter(router *gin.RouterGroup) {
	AdvertApi := api.ApiGroupApp.AdvertApi

	router.POST("/adverts",middleware.JwtAuth(), AdvertApi.AdvertCreateView)
	router.DELETE("/adverts",middleware.JwtAuth(), AdvertApi.AdvertRemoveView)
	router.PUT("/adverts/:id",middleware.JwtAuth(), AdvertApi.AdvertUpdateView)
	router.GET("/adverts",middleware.JwtAuth(), AdvertApi.AdvertListView)

}
