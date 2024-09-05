package images_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)
// ImageListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
func (ImagesApi) ImageListView(c *gin.Context) {

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	logrus.Info(page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//使用封装的列表分页查询服务
	list, count, err := common.ComList(
		models.BannerModel{}, 
		common.Option{
			PageInfo: page, 
			Debug: true,
		})
	if err != nil {
		res.FailWithMessage("图片列表获取失败", c)
		return
	}
	res.OkWithList(list, count, c)
}
