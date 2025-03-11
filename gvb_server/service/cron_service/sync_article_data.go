package cron_service

import (
	"context"
	"encoding/json"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_service"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// 同步文章点赞浏览评论数，从redis到es中
// 总体思路：拿到es中的所有文章数据，更新有变化的文章数据
func SyncArticleData() {
	// 1.查询es中的所有文章数据，为后续的数据更新做准备
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	// 2.拿到redis中的缓存数据，文章点赞浏览收藏数据
	diggInfo := redis_service.NewDigg().GetInfo()
	lookInfo := redis_service.NewArticleLook().GetInfo()
	commentInfo := redis_service.NewCommentCount().GetInfo()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			logrus.Error(err)
			continue
		}
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		// 3.计算新的数据，老的数据加上缓存的数据
		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment
		// 4.判断数据是否变化，如果变化，则更新es中的数据
		if article.DiggCount == newDigg && article.LookCount == newLook && article.CommentCount == newComment {
			logrus.Info(article.Title, " 点赞和浏览以及评论数无变化")
			continue
		}
		// 5.更新es中的数据·	
		_, err = global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count":    newDigg,
				"look_count":    newLook,
				"comment_count": newComment,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Infof("%s, 点赞浏览评论数据同步成功, 点赞数：%d, 浏览数：%d, 评论数 %d", article.Title, newDigg, newLook, newComment)
	}
	// 6.清空redis中的缓存数据
	redis_service.NewDigg().Clear()
	redis_service.NewArticleLook().Clear()
	redis_service.NewCommentCount().Clear()
}
