package models

import (
	"context"
	"encoding/json"
	"gvb_server/global"
	"gvb_server/models/ctype"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// 文章表
type ArticleModel struct {
	ID        string `json:"id"  structs:"id"`                // es的ID
	CreatedAt string `json:"created_at" structs:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at" structs:"updated_at"` // 更新时间

	Title    string `json:"title" structs:"title"`                // 文章标题
	Keyword  string `json:"keyword,omit(list)" structs:"keyword"` //关键字    存储时与文章标题相同，便于搜索         omit排除，filter.Omit("list", list)
	Abstract string `json:"abstract" structs:"abstract"`          // 文章简介  如果不存在文章简介，截取文章内容的100字为简介
	Content  string `json:"content,omit(list)" structs:"content"` // 文章内容

	LookCount     int `json:"look_count" structs:"look_count"`         // 文章浏览量
	CommentCount  int `json:"comment_count" structs:"comment_count"`   // 文章评论量
	DiggCount     int `json:"digg_count" structs:"digg_count"`         // 文章点赞量
	CollectsCount int `json:"collects_count" structs:"collects_count"` // 文章收藏量

	UserID       uint   `json:"user_id" structs:"user_id"`               // 用户id
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"` // 用户昵称
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`       // 用户头像

	Category string `json:"category" structs:"category"`        // 文章分类
	Source   string `json:"source,omit(list)" structs:"source"` // 文章来源
	Link     string `json:"link,omit(list)" structs:"link"`     // 原文链接

	BannerID  uint   `json:"banner_id" structs:"banner_id"`   // 文章封面id
	BannerUrl string `json:"banner_url" structs:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags" structs:"tags"` // 文章标签
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "keyword": { 
        "type": "keyword"
      },
      "abstract": { 
        "type": "text"
      },
      "content": { 
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "keyword"
      },
      "user_avatar": { 
        "type": "keyword"
      },
      "category": { 
        "type": "keyword"
      },
      "source": { 
        "type": "keyword"
      },
      "link": { 
        "type": "keyword"
      },
      "banner_id": {
        "type": "integer"
      },
      "banner_url": { 
        "type": "keyword"
      },
      "tags": { 
        "type": "keyword"
      },
      "created_at":{
        "type": "date",
        "null_value": "1970-01-01 00:00:00",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updated_at":{
        "type": "date",
        "null_value": "1970-01-01 00:00:00",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

// IndexExists 索引是否存在
func (a ArticleModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

// CreateIndex 创建索引,相当于创建表
func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		// 有索引
		a.RemoveIndex()
	}
	// 没有索引
	// 创建索引
	createIndex, err := global.ESClient.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}
	logrus.Infof("索引 %s 创建成功", a.Index())
	return nil
}

// RemoveIndex 删除索引，相当于删除表
func (a ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}

// Create 添加的方法
func (a *ArticleModel) Create() (err error) {
	indexResponse, err := global.ESClient.Index().
		Index(a.Index()).
		BodyJson(a).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// ISExistData 是否存在该文章
func (a ArticleModel) ISExistData() bool {
	res, err := global.ESClient.
		Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

// GetDataById 根据id获取数据
func (a *ArticleModel) GetDataById(id string) error {
	res, err := global.ESClient.
		Get().
		Index(a.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return err
	}
	err = json.Unmarshal(res.Source, a)
	return err
}
