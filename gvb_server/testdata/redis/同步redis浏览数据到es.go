package main

import (
	"context"
	"encoding/json"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_service"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	// 配置信息读取
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()
	global.Redis = core.ConnectRedis()
	global.ESClient = core.EsConnect()

	// 点赞数同步es
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	lookInfo := redis_service.GetLookInfo()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err := json.Unmarshal(hit.Source, &article)

		look := lookInfo[hit.Id]
		newLook := article.LookCount + look
		if article.LookCount == newLook {
			logrus.Info(article.Title, "浏览数无变化")
			continue
		}
		_, err = global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"look_count": newLook,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Info(article.Title, "浏览数据同步成功， 浏览数", newLook)
	}
	redis_service.DiggClear()
}
