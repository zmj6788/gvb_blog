package menu_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}
type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 菜单列表
// @Tags 菜单管理
// @Summary 菜单列表
// @Description 菜单列表
// @Router /api/menus [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]MenuResponse}
func (*MenuApi) MenuListView(c *gin.Context) {
	//先查出所有菜单
	var menuList []models.MenuModel
	//取出菜单id
	var menuIdList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIdList)

	//查连接表
	//用菜单id去查出菜单关联的图片数据
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIdList)
	//封装响应数据
	var menus []MenuResponse
	for _, menu := range menuList {
		//menu 就是一个菜单

		//遍历菜单关联的图片数据
		// var banners []Banner  只声明，不赋值，如果是引用类型，那么最后就等于 nil，在前端表示就是 null
		//修改后解决这个问题
		var banners = make([]Banner, 0)
		for _, banner := range menuBanners {
			if menu.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}

	res.OkWithData(menus, c)
}
