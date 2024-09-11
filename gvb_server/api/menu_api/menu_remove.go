package menu_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MenuRemoveView 批量删除菜单
// @Tags 菜单管理
// @Summary 批量删除菜单
// @Description 批量删除菜单
// @Param data body models.RemoveRequest    true  "菜单id列表"
// @Router /api/menus [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (*MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	// 判断菜单是否存在
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}

	// 删除菜单前先删除菜单和banner的关联
	// 然后才能删除菜单
	// 使用事务
	// 确保整个操作（清除关联和删除菜单）都在事务内处理，
	// 并且在操作过程中出现任何问题都会导致事务回滚，从而保持数据的完整性。

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 使用事务对象来清除关联
		if err := tx.Model(&menuList).Association("Banners").Clear(); err != nil {
			global.Log.Error(err)
			return err
		}

		// 使用事务对象来删除 menuList
		if err := tx.Delete(&menuList).Error; err != nil {
			global.Log.Error(err)
			return err
		}

		return nil
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}

	res.OkWithMessage(fmt.Sprintf("共删除 %d 个菜单", count), c)
}
