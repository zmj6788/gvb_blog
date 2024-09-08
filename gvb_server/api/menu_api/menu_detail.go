package menu_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)
// MenuDetailView 菜单详情
// @Tags 菜单管理
// @Summary 菜单详情
// @Description 菜单详情
// @Param id path int true "id"
// @Router /api/menus/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=MenuResponse}
func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	// 先获取菜单数据
	var menuModel models.MenuModel
	err := global.DB.First(&menuModel, id).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单不存在", c)
		return
	}
	
	// 获取菜单的图片关联
	var menuBanners []models.MenuBannerModel
	//Preload("BannerModel")：这是预加载 BannerModel 关联的数据，以避免 N+1 查询问题。
	// 这意味着在查询 menuBanners 的同时，会预加载每个 menuBanner 的 BannerModel 关联数据。
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", menuModel.ID)
	
	var banners = make([]Banner, 0)
	for _, banner := range menuBanners {
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	// 获取菜单详情
	menus := MenuResponse{menuModel, banners}
	res.OkWithData(menus, c)

}
