package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRouter(router *gin.RouterGroup) {
	CommentApi := api.ApiGroupApp.CommentApi

	router.POST("/comments", middleware.JwtAuth(), CommentApi.CommentCreateView)
	router.GET("/comments", CommentApi.CommentListView)

}
