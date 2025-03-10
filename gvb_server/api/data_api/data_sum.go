package data_api

import (
	"context"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type DataSumResponse struct {
	UserCount      int `json:"user_count"`  // 用户总数
	ArticleCount   int `json:"article_count"` // 文章总数
	MessageCount   int `json:"message_count"` // 消息总数
	ChatGroupCount int `json:"chat_group_count"` // 群组消息总数
	NowLoginCount  int `json:"now_login_count"` // 今日登录总数
	NowSignCount   int `json:"now_sign_count"` // 今日注册总数
}
// DataSumView 一些总数的统计
// @Tags 数据管理
// @Summary 一些总数的统计
// @Description 一些总数的统计
// @Router /api/data_sum [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (DataApi) DataSumView(c *gin.Context) {

	var userCount, articleCount, messageCount, ChatGroupCount int
	var nowLoginCount, nowSignCount int

	result, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())
	articleCount = int(result.Hits.TotalHits.Value) //搜索到结果总条数
	global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&userCount)
	global.DB.Model(models.MessageModel{}).Select("count(id)").Scan(&messageCount)
	global.DB.Model(models.ChatModel{IsGroup: true}).Select("count(id)").Scan(&ChatGroupCount)
	global.DB.Model(models.LoginDataModel{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowLoginCount)
	global.DB.Model(models.UserModel{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowSignCount)

	res.OkWithData(DataSumResponse{
		UserCount:      userCount,
		ArticleCount:   articleCount,
		MessageCount:   messageCount,
		ChatGroupCount: ChatGroupCount,
		NowLoginCount:  nowLoginCount,
		NowSignCount:   nowSignCount,
	}, c)
}
