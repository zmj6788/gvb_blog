package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"

	"github.com/olivere/elastic/v7"
)

type JSONData struct {
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

type Response struct {
	Tags         string    `json:"tags"`
	ArticleCount int    `json:"article_count"`
	ArticleList  []Article `json:"article_list"`
}
type Article struct {
	Title string `json:"title"`
}

func init() {
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()

	//es连接
	global.ESClient = core.EsConnect()

}

func main() {
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
	// agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("_id"))

	query := elastic.NewBoolQuery()

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())

	if err != nil {
		global.Log.Error(err.Error())
		return
	}

	var jsonData JSONData
	_ = json.Unmarshal(result.Aggregations["tags"], &jsonData)


	// 最终输出结果
	var reslist = make([]Response,0) 
	for _, bucket := range jsonData.Buckets {
		var res Response
		res.Tags = bucket.Key
		res.ArticleCount = bucket.DocCount
		for _, bucket2 := range bucket.Articles.Buckets {
			var article Article
			article.Title = bucket2.Key
			res.ArticleList = append(res.ArticleList, article)
		}
		reslist = append(reslist, res)
	}
	// 输出结果03
	// [{node 2 [{node基础} {node高级}]} {前端 2 [{html基础} {vue基础}]}
	//  {go 1 [{go基础}]} {html 1 [{html基础}]} {vue 1 [{vue基础}]}]
	// json格式
	// {
    // "code": 200,
    // "data": {
    //     "count": 5,
    //     "list": [
    //         {
    //             "tags": "node",
    //             "article_count": 2,
    //             "article_list": [
    //                 {
    //                     "title": "node基础"
    //                 },
    //                 {
    //                     "title": "node高级"
    //                 }
    //             ]
    //         },
		// ...
    //  ]
    // }
	// }
	// 符合需求
	fmt.Println(reslist)
}

/*
{
    "doc_count_error_upper_bound": 0,
    "sum_other_doc_count": 0,
    "buckets": [
        {
            "key": "node",
            "doc_count": 2,
            "articles": {
                "doc_count_error_upper_bound": 0,
                "sum_other_doc_count": 0,
                "buckets": [
                    {
                        "key": "node基础",
                        "doc_count": 1
                    },
                    {
                        "key": "node高级",
                        "doc_count": 1
                    }
                ]
            }
        }
    ]
}


*/
