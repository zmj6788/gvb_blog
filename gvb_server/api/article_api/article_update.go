package article_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/es_service"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ArticleUpdateRequest struct {
	Title    string   `json:"title" `    // 文章标题
	Abstract string   `json:"abstract" ` // 文章简介
	Content  string   `json:"content" `  // 文章内容
	Category string   `json:"category"`  // 文章分类
	Source   string   `json:"source"`    // 文章来源
	Link     string   `json:"link"`      // 原文链接
	BannerID uint     `json:"banner_id"` // 文章封面id
	Tags     []string `json:"tags"`      // 文章标签
	ID       string   `json:"id" `
}

// ArticleUpdateView 更新文章
// @Tags 文章管理
// @Summary 更新文章
// @Description 更新文章
// @Param data body ArticleUpdateRequest    true  "广告的一些参数"
// @Router /api/articles [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}

	// 判断文章封面图片是否存在
	var bannerUrl string
	if cr.BannerID != 0 {
		err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
		if err != nil {
			res.FailWithMessage("banner不存在", c)
			return
		}
	}
	article := models.ArticleModel{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: bannerUrl,
		Tags:      cr.Tags,
	}

	// 注意使用structs是的时候，一定要加 structs标签，不然转换之后的map就是大写的key
	// structs标签加到结构体上
	maps := structs.Map(&article)
	var DataMap = map[string]any{}
	// 去掉空值
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}
	err = article.GetDataById(cr.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("文章不存在", c)
		return
	}

	// Doc(DataMap)时，Elasticsearch客户端库负责将这个map转换为JSON格式
	// json更新es数据
	err = es_service.ArticleUpdate(cr.ID, DataMap)
	if err != nil {
		logrus.Error(err.Error())
		res.FailWithMessage("更新失败", c)
		return
	}
	// 文章更新后，需要同步文章数据到全文搜索的表中
	newArticle, _ := es_service.CommDetail(cr.ID)
	if article.Content != newArticle.Content || article.Title != newArticle.Title {
		// 删除原有的
		es_service.DeleteFullTextByArticleID(cr.ID)
		// 添加现在的
		es_service.AsyncArticleByFullText(cr.ID, newArticle.Title, newArticle.Content)
	}
	res.OkWithMessage("更新成功", c)
}
