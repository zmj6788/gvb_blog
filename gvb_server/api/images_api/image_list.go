package images_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// 用于分页查询
type Page struct {
	Page  int    `form:"page"`  //页码
	Key   string `form:"key"`   //搜索关键字
	Limit int    `form:"limit"` //每页显示多少条
	Sort  string `form:"sort"`  //排序
}

func (ImagesApi) ImageListView(c *gin.Context) {

	var page Page
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imagesList []models.BannerModel
	count := global.DB.Find(&imagesList).RowsAffected
	//偏移量
	offset := (page.Page - 1) * page.Limit
	//如果偏移量小于0，则从0开始
	if offset < 0 {
		offset = 0
	}
	//如果limit为0，则查询所有
	if page.Limit == 0 {
		page.Limit = -1
	}
	err = global.DB.Limit(page.Limit).Offset(offset).Find(&imagesList).Error
	if err != nil {
		res.FailWithMessage("图片列表获取失败", c)
		return
	}
	res.OkWithData(gin.H{"count": count, "list": imagesList}, c)
}
