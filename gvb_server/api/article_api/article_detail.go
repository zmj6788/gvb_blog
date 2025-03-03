package article_api

import (
	"gvb_server/models/res"
	"gvb_server/service/es_service"
	"gvb_server/service/redis_service"

	"github.com/gin-gonic/gin"
)

type ArticleDetailRequest struct {
	ID string `uri:"id" json:"id" form:"id"`
}

// ArticleDetailView 文章详情
// @Tags 文章管理
// @Summary 文章详情
// @Description 文章详情
// @Param id path string true "id"
// @Router /api/articles/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleDetailView(c *gin.Context) {

	var cr ArticleDetailRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	
	// 查看文章详情时，浏览量增加
	redis_service.NewArticleLook().Set(cr.ID)
	model, err := es_service.CommDetail(cr.ID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}

	res.OkWithData(model, c)

}

type ArticleDetailByTitleRequest struct {
	Title string `json:"title" form:"title"`
}

// ArticleDetailByTitleView 文章详情
// @Tags 文章管理
// @Summary 文章详情
// @Description 文章详情
// @Router /api/articles/detail [get]
// @Param title query string true "title"
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {

	var cr ArticleDetailByTitleRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	
	model, err := es_service.CommDetailByKeyword(cr.Title)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	res.OkWithData(model, c)

}
