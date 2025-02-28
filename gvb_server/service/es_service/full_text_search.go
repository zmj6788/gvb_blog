package es_service

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
)

type SearchData struct {
	Key string `json:"key"`     //关联文章id，便于删除
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n")
	var isCode bool = false
	var headList, bodyList []string
	var body string
	// 文章标题
	headList = append(headList, getHeader(title))
	for _, s := range dataList {
		// #{1,6}
		// 判断一下是否是代码块
		if strings.HasPrefix(s, "```") {
			// 排除代码块中带有#的情况
			isCode = !isCode
		}
		if strings.HasPrefix(s, "#") && !isCode {
			headList = append(headList, getHeader(s))
			// 如果清除空行，文章标题下没有不带标题的内容，无法正常处理
			//if strings.TrimSpace(body) != "" {
			// 到达下一个标题是加入上一个标题的正文到bodyList中
			bodyList = append(bodyList, getBody(body))
			//}
			// 文章标题下的正文为空情况，是如何来的，从这里来的
			body = ""
			continue
		}
		// 不是标题的被赋值为body，将一个标题下的所有正文拼接加入bodyList中
		body += s
	}
	// 将最后一个标题的正文加入bodyList中
	bodyList = append(bodyList, getBody(body))

	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
			Key:  id,
		})
	}
	b, _ := json.Marshal(searchDataList)
	fmt.Println(string(b))
	// fmt.Println(len(headList), len(bodyList))
	return searchDataList
}

// 标题格式处理
func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}

// 正文格式处理
func getBody(body string) string {
	// 将markdown转为html
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	// 从html中获取文本内容
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

// 跳转链接格式处理
func getSlug(slug string) string {
	return "#" + slug
}

// 同步一个文章数据到全文搜索的表中
func AsyncArticleByFullText(id, title, content string)  {
	// 处理文章数据为支持全文搜索的格式
	searchDataList := GetSearchIndexDataByContent(id, title, content)
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
		return 
	}
	logrus.Infof("%s 添加成功,共 %d 条",title, len(result.Succeeded()))
}

// 删除文章在全文搜索中的数据
func DeleteFullTextByArticleID(id string) {
  boolSearch := elastic.NewTermQuery("key", id)
  res, _ := global.ESClient.
    DeleteByQuery().
    Index(models.FullTextModel{}.Index()).
    Query(boolSearch).
    Do(context.Background())
  logrus.Infof("成功删除 %d 条记录", res.Deleted)
}