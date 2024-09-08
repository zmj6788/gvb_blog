package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func MenuRouter(router *gin.RouterGroup) {
	MenuApi := api.ApiGroupApp.MenuApi
	router.GET("/menus", MenuApi.MenuListView)
	router.GET("/menu_names", MenuApi.MenuNameListView)
	router.GET("/menus/:id", MenuApi.MenuDetailView)
	router.POST("/menus", MenuApi.MenuCreateView)
	router.PUT("/menus/:id", MenuApi.MenuUpdateView)
	router.DELETE("/menus", MenuApi.MenuRemoveView)
}
