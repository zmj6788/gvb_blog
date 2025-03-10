package article_api

import (
	"context"
	"encoding/json"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)
// FullTextSearchView 全文搜索
// @Tags 文章管理
// @Summary 全文搜索
// @Description 全文搜索
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/articles/text [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.FullTextModel]}
func (ArticleApi) FullTextSearchView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	boolQuery := elastic.NewBoolQuery()

	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}

	result, err := global.ESClient.
		Search(models.FullTextModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")).
		Size(100).
		Do(context.Background())
	if err != nil {
		return
	}
	count := result.Hits.TotalHits.Value //搜索到结果总条数
	fullTextList := make([]models.FullTextModel, 0)
	for _, hit := range result.Hits.Hits {
		var model models.FullTextModel
		json.Unmarshal(hit.Source, &model)

		body, ok := hit.Highlight["body"]
		if ok {
			model.Body = body[0]
		}

		fullTextList = append(fullTextList, model)
	}

	res.OkWithList(fullTextList, count, c)
}
