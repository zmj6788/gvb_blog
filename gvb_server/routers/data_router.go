package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func DataRouter(router *gin.RouterGroup) {
	DataApi := api.ApiGroupApp.DataApi
	router.GET("/data_seven_login", DataApi.SevenLogin)
	router.GET("/data_sum", DataApi.DataSumView)

}
