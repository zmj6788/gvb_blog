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
	Tags         string    `json:"tags"`
	ArticleCount int       `json:"article_count"`
	ArticleIDList  []string `json:"article_id_list"`
}


func (ArticleApi) ArticleTagListView(c *gin.Context) {

	// 期望结果
	// [{"tags": "node","article_count": 2, "article_list": [{},{}]"}]

	// 聚合，用来对"tags"字段进行分组计数，tags相同的分为一组
	agg := elastic.NewTermsAggregation().Field("tags")

	// 输出结果01
	//{"doc_count_error_upper_bound":0,"sum_other_doc_count":0,
	// "buckets":[{"key":"node","doc_count":2},{"key":"go","doc_count":1}]}

	// 添加了一个子聚合"articles"，tags相同的分组内，再对"keyword"字段进行分组，
	// 输出结果02
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))

	query := elastic.NewBoolQuery()

	result, err := global.ESClient.
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

	// 最终输出结果
	var reslist = make([]ArticleTagResponse,0) 
	for _, bucket := range jsonData.Buckets {
		var atres ArticleTagResponse
		atres.Tags = bucket.Key
		atres.ArticleCount = bucket.DocCount
		for _, bucket2 := range bucket.Articles.Buckets {
			atres.ArticleIDList = append(atres.ArticleIDList, bucket2.Key)
		}
		reslist = append(reslist, atres)
	}
	res.OkWithList(reslist, result.Hits.TotalHits.Value, c)
}
