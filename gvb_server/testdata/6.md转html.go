package main

import (
	"fmt"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

func main() {
	//md转html
	unsafe := blackfriday.MarkdownCommon([]byte("### 你好\n ```python\nprint('你好')\n```\n - 123 \n \n<script>alert(123)</script>\n\n ![图片](http://xxx.com)"))
	fmt.Println(string(unsafe))

	//html获取文本内容，xss过滤
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	// fmt.Println(doc.Text())
	doc.Find("script").Remove()  // 删除script标签以及里面内容
	fmt.Println(doc.Text())

	//html转md
	converter := md.NewConverter("", true, nil)
	html, _ := doc.Html()
	markdown, err := converter.ConvertString(html)
	fmt.Println(markdown, err)
}
