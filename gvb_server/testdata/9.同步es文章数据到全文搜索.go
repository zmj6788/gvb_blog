package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/es_service"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	// 配置信息读取
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()
	//es连接
	global.ESClient = core.EsConnect()

	// 同步es文章数据到全文搜索的表中
	query := elastic.NewMatchAllQuery()

	result, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Size(1000).
		Do(context.Background())

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)
		// 处理文章数据为支持全文搜索的格式
		searchDataList := es_service.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)
		// fmt.Println(searchDataList)
		// 将数据入库，es库
		// 批量添加数据到es
		bulk := global.ESClient.Bulk()
		for _, indexData := range searchDataList {
			req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
			bulk.Add(req)
		}
		result, err := bulk.Do(context.Background())
		if err != nil {
			logrus.Error(err)
			continue
		}
		fmt.Println(article.Title, "添加成功", "共", len(result.Succeeded()), " 条！")
	}

}
