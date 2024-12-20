package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"

	"github.com/gin-gonic/gin"
)

func ArticleRouter(router *gin.RouterGroup) {
	ArticleApi := api.ApiGroupApp.ArticleApi

	router.POST("/articles", middleware.JwtAuth(), ArticleApi.ArticleCreateView)
	
	router.GET("/articles", ArticleApi.ArticleListView)
	router.GET("/articles/:id", ArticleApi.ArticleDetailView)
	router.GET("/articles/detail", ArticleApi.ArticleDetailByTitleView)
	router.GET("/articles/calendar", ArticleApi.ArticleCalendarView)
	router.GET("/articles/tags", ArticleApi.ArticleTagListView)

	router.PUT("/articles", ArticleApi.ArticleUpdateView)
	router.DELETE("/articles", ArticleApi.ArticleRemoveView)
}
