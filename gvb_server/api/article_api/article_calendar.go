package article_api

import (
	"context"
	"encoding/json"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type CalendarResponse struct {
	Date string `json:"date"`
	Count int `json:"count"`
}

var DateCount = map[string]int{}

type JSONData struct {
		Buckets []struct {
			KeyAsString string `json:"key_as_string"`
			Key int64 `json:"key"`
			DocCount int `json:"doc_count"`
		} `json:"buckets"`
}
// ArticleCalendarView 文章日历
// @Tags 文章管理
// @Summary 文章日历
// @Description 文章日历
// @Router /api/articles/calendar [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[CalendarResponse]}
func (ArticleApi) ArticleCalendarView(c *gin.Context) {

	//时间聚合

	agg := elastic.NewDateHistogramAggregation().
		Field("created_at").
		CalendarInterval("day")

	//时间段搜索
	//从今天开始到去年的今天

	format := "2006-01-02 15:04:05"
	now := time.Now()
	aYearAgo := now.AddDate(-1, 0, 0)

	// gte: 大于等于  lte: 小于等于

	query := elastic.NewRangeQuery("created_at").Gte(aYearAgo.Format(format)).Lte(now.Format(format))

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())

	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("查询失败", c)
		return
	}

	//获取到一年的数据后解析到go结构体中
	var jsonData JSONData
	_ = json.Unmarshal(result.Aggregations["calendar"],&jsonData)

	//将每天的数据存到map中，便于高效查询

	for _, v := range jsonData.Buckets {
		//时间格式转换
		Time, _ := time.Parse(format, v.KeyAsString)
		DateCount[Time.Format("2006-01-02")] = v.DocCount
	}

	//将每日数据遍历添加到响应列表中
	//避免出现null

	var reslist = make([]CalendarResponse, 0)
	days := int(now.Sub(aYearAgo).Hours() / 24)
	for i := 0; i <= days; i++ {
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")
		//查询不到默认为0
		count, _ := DateCount[day]
		reslist = append(reslist, CalendarResponse{
			Date: day,
			Count: count,
		})
	}

	res.OkWithData(reslist, c)

}
