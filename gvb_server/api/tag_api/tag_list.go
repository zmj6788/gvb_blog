package tag_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"

	"github.com/gin-gonic/gin"
)

// TagListView 标签列表
// @Tags 标签管理
// @Summary 标签列表
// @Description 标签列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/tags [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.TagModel]}
func (TagApi) TagListView(c *gin.Context) {

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//使用封装的列表分页查询服务
	list, count, err := common.ComList(
		models.TagModel{},
		common.Option{
			PageInfo: page,
			Debug:    true,
		})
	if err != nil {
		res.FailWithMessage("标签列表获取失败", c)
		return
	}
	
	//需要显示这个标签下关联文章的数量
	
	res.OkWithList(list, count, c)
}
