package advert_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (*AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advert models.AdvertModel
	err = global.DB.First(&advert, id).Error
	if err != nil {
		res.FailWithMessage("广告不存在", c)
		return
	}
	// 结构体转map的第三方包
	// 用map原因，让布尔值可以正常修改
	// 删除binding:"required" msg:"请选择是否显示" 
	maps := structs.Map(&cr)
	fmt.Println(maps)
	err = global.DB.Model(&advert).Updates(maps).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败", c)
		return
	}

	res.OkWithMessage("修改广告成功", c)
}
