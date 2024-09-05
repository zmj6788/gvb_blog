package advert_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"strings"

	"github.com/gin-gonic/gin"
)

func (*AdvertApi) AdvertListView(c *gin.Context) {

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//分为两种情况，前台只需要显示的广告数据，后台需要显示所有数据
	//使用referer判断请求是从哪里发出的
	//实现差异化数据展示

	referer := c.GetHeader("Referer")
	//只查显示的广告数据
	is_show := true
	if strings.Contains(referer, "admin"){
		//查全部
		is_show = false 
	}

	//使用封装的列表分页查询服务
	list, count, err := common.ComList(
		models.AdvertModel{IsShow: is_show},
		common.Option{
			PageInfo: page,
			Debug:    true,
		})
	if err != nil {
		res.FailWithMessage("广告列表获取失败", c)
		return
	}
	res.OkWithList(list, count, c)
}
