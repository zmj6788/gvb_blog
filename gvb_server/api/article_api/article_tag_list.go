package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type TagsJSONData struct {
	Buckets []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

type ArticleTagResponse struct {
	Tag           string   `json:"tag"`
	ArticleCount  int      `json:"article_count"`
	ArticleIDList []string `json:"article_id_list"`
	CreatedAt     string   `json:"created_at"` // 标签的创建时间
}

// ArticleTagListView 文章标签列表
// @Tags 文章管理
// @Summary 文章标签列表
// @Description 文章标签列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/articles/tags [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[ArticleTagResponse]}
func (ArticleApi) ArticleTagListView(c *gin.Context) {

	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	if cr.Limit == 0 {
		cr.Limit = 10
	}
	if cr.Page == 0 {
		cr.Page = 1
	}
	// 获取偏移量,从第几条开始查询
	offset := (cr.Page - 1) * cr.Limit
	// 文章列表接口
	// 1.聚合查询
	// 2.聚合分页
	// 3.聚合后的总数获取
	// 4.同步mysql数据库，从mysql的tagsmodels中取出create_at给到我们对应标签列表的响应中

	// 3.聚合后的总数获取,其实直接也可以获取总数result.Hits.TotalHits.Value
	// 下方代码不需要
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewCardinalityAggregation().Field("tags")).
		Size(0).
		Do(context.Background())
	cTag, _ := result.Aggregations.Cardinality("tags")
	count := int64(*cTag.Value)

	// 1.聚合查询
	// 聚合，用来对"tags"字段进行分组计数，tags相同的分为一组
	agg := elastic.NewTermsAggregation().Field("tags")

	// 添加了一个子聚合"articles"，tags相同的分组内，再对"keyword"字段进行分组，
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))

	// 2.聚合分页
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))

	query := elastic.NewBoolQuery()

	result, err = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())

	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("查询文章标签列表失败", c)
		return
	}

	var jsonData TagsJSONData
	_ = json.Unmarshal(result.Aggregations["tags"], &jsonData)

	// 解决response.CreatedAt赋值不成功问题，使用指针
	var reslist = make([]*ArticleTagResponse, 0)
	// 存储所有es中查询到的标签
	var tagStringList []string
	for _, bucket := range jsonData.Buckets {
		var atres ArticleTagResponse
		atres.Tag = bucket.Key
		atres.ArticleCount = bucket.DocCount
		for _, bucket2 := range bucket.Articles.Buckets {
			atres.ArticleIDList = append(atres.ArticleIDList, bucket2.Key)
		}
		reslist = append(reslist, &atres)
		tagStringList = append(tagStringList, bucket.Key)
	}

	// 4.同步mysql数据库
	var tagModelList []models.TagModel
	global.DB.Find(&tagModelList, "title in ?", tagStringList)
	var tagDate = map[string]string{}
	for _, model := range tagModelList {
		tagDate[model.Title] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}
	fmt.Println("tagDate", tagDate)
	for _, response := range reslist {
		response.CreatedAt = tagDate[response.Tag]
	}
	res.OkWithList(reslist, count, c)
}
