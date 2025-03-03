package es_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_service"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// CommList 列表查询
// CommList 升级版本，搜索内容可以自主增加。标题高亮.可以支持排序
func CommList(o Option) (List []models.ArticleModel, count int64, err error) {
	boolSearch := elastic.NewBoolQuery()
	if o.Key != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(o.Key, o.Fields...),
		)
	}

	if o.Tag != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(o.Tag, "tags"),
		)
	}

	sortField := SortField{
		Field:     "created_at",
		Ascending: false, // 从小到大  从大到小
	}
	// 自定义排序
	if o.Sort != "" {
		// 按照 空格 分割字符串
		_list := strings.Split(o.Sort, " ")
		if len(_list) == 2 && (_list[1] == "desc" || _list[1] == "asc") {
			sortField.Field = _list[0]
			if _list[1] == "desc" {
				// 默认降序排序
				sortField.Ascending = false
			}
			if _list[1] == "asc" {
				// 时间升序排列
				sortField.Ascending = true
			}
		}
	}

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Highlight(elastic.NewHighlight().Field("title")).
		From(o.GetForm()).
		Sort(sortField.Field, sortField.Ascending).
		Size(o.Limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count = int64(res.Hits.TotalHits.Value) //搜索到结果总条数
	diggInfo := redis_service.NewDigg().GetInfo()
	lookInfo := redis_service.NewArticleLook().GetInfo()
	commentInfo := redis_service.NewCommentCount().GetInfo()
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

		// 显示时应用高亮
		title, ok := hit.Highlight["title"]
		if ok {
			fmt.Println(title)
			model.Title = title[0]
		}
		model.ID = hit.Id
		// 更新点赞数,但是没有同步数据到es
		// 没有同步但仍然能正确显示点赞数了，因为redis里的数据没有被清空
		digg := diggInfo[hit.Id]
		model.DiggCount = model.DiggCount + digg
		look := lookInfo[hit.Id]
		model.LookCount = model.LookCount + look
		comment := commentInfo[hit.Id]
		model.CommentCount = model.CommentCount + comment
		List = append(List, model)
	}
	return List, count, err
}

// id查询
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
	model.LookCount = model.LookCount + redis_service.NewArticleLook().Get(id)
	return
}

// 根据keyword查询
func CommDetailByKeyword(key string) (model models.ArticleModel, err error) {

	res, err := global.ESClient.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return model, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]

	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		return
	}
	model.ID = hit.Id
	redis_service.NewArticleLook().Set(hit.Id)
	model.LookCount = model.LookCount + redis_service.NewArticleLook().Get(hit.Id)
	return
}

func ArticleUpdate(id string, data map[string]any) error {
	_, err := global.ESClient.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(data).
		Do(context.Background())
	return err
}
