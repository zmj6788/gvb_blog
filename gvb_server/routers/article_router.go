package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"

	"github.com/gin-gonic/gin"
)

func ArticleRouter(router *gin.RouterGroup) {
	ArticleApi := api.ApiGroupApp.ArticleApi

	router.POST("/articles", middleware.JwtAuth(),ArticleApi.ArticleCreateView)
	// router.DELETE("/articles", ArticleApi.AdvertRemoveView)
	// router.PUT("/articles/:id", ArticleApi.AdvertUpdateView)
	router.GET("/articles", ArticleApi.ArticleListView)

}
