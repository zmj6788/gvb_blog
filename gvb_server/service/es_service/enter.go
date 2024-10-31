package es_service

import (
	"context"
	"encoding/json"
	"gvb_server/global"
	"gvb_server/models"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// CommList 列表查询
func CommList(key string, page, limit int) (List []models.ArticleModel, count int64, err error) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count = int64(res.Hits.TotalHits.Value) //搜索到结果总条数
	//将es中的数据解析到go结构体中
	for _, hit := range res.Hits.Hits {
		var model models.ArticleModel
		// 将 hit.Source 对象序列化为 JSON 格式的字节切片 data
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		// 将 datajson数据 反序列化为 ArticleModel对象
		err = json.Unmarshal(data, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		model.ID = hit.Id
		List = append(List, model)
	}
	return List, count,err
}

func CommDetail(id string) (model models.ArticleModel, err error) {
	
	res, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return
	}
	data, err := res.Source.MarshalJSON()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &model)
	model.ID = res.Id
	return
}