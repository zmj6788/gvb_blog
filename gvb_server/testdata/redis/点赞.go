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

	err := redis_service.Digg("2DtL-5IBBOEDMw_pwTmW")
	if err != nil {
		global.Log.Error(err.Error())
	}
	global.Log.Info("点赞成功")

	// global.Log.Info(redis_service.GetDigg("1ztL-5IBBOEDMw_pRTnk"))
	global.Log.Info(redis_service.GetDiggInfo())
	// redis_service.DiggClear()

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

	diggInfo := redis_service.GetDiggInfo()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err := json.Unmarshal(hit.Source, &article)

		digg := diggInfo[hit.Id]
		newDigg := article.DiggCount + digg
		if article.DiggCount == newDigg {
			logrus.Info(article.Title, "点赞数无变化")
			continue
		}
		_, err = global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count": newDigg,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Info(article.Title, "点赞数据同步成功， 点赞数", newDigg)
	}
	redis_service.DiggClear()
}
