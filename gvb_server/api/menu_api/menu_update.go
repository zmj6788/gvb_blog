package menu_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// MenuUpdateView 更新菜单
// @Tags 菜单管理
// @Summary 更新菜单
// @Param token header string  true  "token"
// @Description 更新菜单
// @Param data body MenuRequest    true  "菜单的一些参数"
// @Param id path int true "id"
// @Router /api/menus/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (MenuApi) MenuUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	// 查询菜单是否存在
	var menu models.MenuModel
	err = global.DB.First(&menu, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 结构体转map的第三方包
	// 用map原因，让布尔值可以正常修改
	// 删除binding:"required" msg:"请选择是否显示"

	//更新MenuModel表数据
	maps := structs.Map(&cr)
	fmt.Println(maps)
	err = global.DB.Model(&menu).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}

	
	// 先清空原来第三张表
	global.DB.Model(&menu).Association("Banners").Clear()
	
	// 给第三张表入库
	if len(cr.ImageSortList) > 0 {
		var menuBannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			// 这里也得判断image_id是否真正有这张图片
			menuBannerList = append(menuBannerList, models.MenuBannerModel{
				MenuID:   menu.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&menuBannerList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("菜单图片关联失败", c)
			return
		}
	}

	res.OkWithMessage("修改菜单成功", c)

}
