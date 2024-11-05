package article_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_service"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)
// ArticleListView 文章列表
// @Tags 文章管理
// @Summary 文章列表
// @Description 文章列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/articles [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.ArticleModel]}
func (ArticleApi) ArticleListView(c *gin.Context) {

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//使用封装的列表分页查询服务
	list, count, err := es_service.CommList(page.Key, page.Page, page.Limit)
	if err != nil {
		res.FailWithMessage("文章列表获取失败", c)
		return
	}
	//json-filter空值问题解决
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OkWithList(list, count, c)
		return
	}
	
	//使用json-filter包，排除某些字段
	res.OkWithList(data, count, c)
}
