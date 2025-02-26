package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

func main() {
	var data = "## 环境搭建\n\n拉取镜像\n\n```Python\ndocker pull elasticsearch:7.12.0\n```\n\n\n\n创建docker容器挂在的目录：\n\n```Python\nmkdir -p /opt/elasticsearch/config & mkdir -p /opt/elasticsearch/data & mkdir -p /opt/elasticsearch/plugins\n\nchmod 777 /opt/elasticsearch/data\n\n```\n\n配置文件\n\n```Python\necho \"http.host: 0.0.0.0\" >> /opt/elasticsearch/config/elasticsearch.yml\n```\n\n\n\n创建容器\n\n```Python\n# linux\ndocker run --name es -p 9200:9200  -p 9300:9300 -e \"discovery.type=single-node\" -e ES_JAVA_OPTS=\"-Xms84m -Xmx512m\" -v /opt/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /opt/elasticsearch/data:/usr/share/elasticsearch/data -v /opt/elasticsearch/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0\n```\n\n\n\n访问ip:9200能看到东西\n\n![](http://python.fengfengzhidao.com/pic/20230129212040.png)\n\n就说明安装成功了\n\n\n\n浏览器可以下载一个 `Multi Elasticsearch Head` es插件\n\n\n\n第三方库\n\n```Go\ngithub.com/olivere/elastic/v7\n```\n\n## es连接\n\n```Go\nfunc EsConnect() *elastic.Client  {\n  var err error\n  sniffOpt := elastic.SetSniff(false)\n  host := \"http://127.0.0.1:9200\"\n  c, err := elastic.NewClient(\n    elastic.SetURL(host),\n    sniffOpt,\n    elastic.SetBasicAuth(\"\", \"\"),\n  )\n  if err != nil {\n    logrus.Fatalf(\"es连接失败 %s\", err.Error())\n  }\n  return c\n}\n```"
	GetSearchIndexDataByContent("/article/hd893bxGHD84", "es的环境搭建", data)
}

type SearchData struct {
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
		})
	}
	b, _ := json.Marshal(searchDataList)
	fmt.Println(string(b))
	fmt.Println(len(headList),len(bodyList))
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
