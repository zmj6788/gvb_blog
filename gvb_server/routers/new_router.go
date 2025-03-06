package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.RouterGroup) {
	NewApi := api.ApiGroupApp.NewApi
	router.GET("/news", NewApi.NewListView)

}
