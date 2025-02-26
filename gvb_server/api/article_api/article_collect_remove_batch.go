package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_service"
	"gvb_server/untils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)
// ArticleCollectBatchRemoveView批量删除我的收藏
// @Tags 文章管理
// @Summary 批量删除我的收藏文章
// @Description 批量删除我的收藏文章
// @Param token header string true "用户token"
// @Param data body models.ESIDListRequest    true  "文章id列表"
// @Router /api/articles/collects [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleCollectBatchRemoveView(c *gin.Context) {
	var cr models.ESIDListRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)

		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var collects []models.UserCollectModel
	var articleIDList []string
	global.DB.Find(&collects, "user_id = ? and article_id in ?", claims.UserID, cr.IDList).
		Select("article_id").
		Scan(&articleIDList)
	if len(articleIDList) == 0 {
		res.FailWithMessage("请求非法", c)
		return
	}
	var idList []interface{}
	for _, s := range articleIDList {
		idList = append(idList, s)
	}
	// 更新文章数
	boolSearch := elastic.NewTermsQuery("_id", idList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	// 更新文章的收藏数
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		count := article.CollectsCount - 1
		err = es_service.ArticleUpdate(hit.Id, map[string]any{
			"collects_count": count,
		})
		if err != nil {
			global.Log.Error(err)
			continue
		}
	}
	// 删除我的收藏记录
	global.DB.Delete(&collects)
	res.OkWithMessage(fmt.Sprintf("成功取消收藏 %d 篇文章", len(articleIDList)), c)

}
