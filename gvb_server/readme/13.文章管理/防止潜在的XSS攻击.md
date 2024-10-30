

```Go
// 校验content  xss
  // 处理content
  // 将markdown转为html
  unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
  // 是不是有script标签
  doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
  //fmt.Println(doc.Text())
  nodes := doc.Find("script").Nodes
  if len(nodes) > 0 {
    // 有script标签,移除
    doc.Find("script").Remove()
    // 将html转为md
    converter := md.NewConverter("", true, nil)
    html, _ := doc.Html()
    markdown, _ := converter.ConvertString(html)
    cr.Content = markdown
  }
```



这段代码的目的是为了检查并移除Markdown内容中的`<script>`标签，以防止潜在的XSS攻击。具体步骤如下：

1. **将Markdown转换为HTML**：使用`blackfriday`库将Markdown格式的内容转换成HTML。
2. **检查是否存在`<script>`标签**：使用`goquery`库解析生成的HTML，并查找所有的`<script>`标签。
3. **移除`<script>`标签**：如果找到了`<script>`标签，则从文档中移除它们。
4. **将处理后的HTML重新转换回Markdown**：使用`github.com/yuin/goldmark`库将清理后的HTML再转换回Markdown格式。

### 为什么需要来回转换？

- **安全检查**：直接在Markdown文本中查找和移除`<script>`标签可能不够可靠，因为Markdown语法并不直接支持`<script>`标签，而是在转换为HTML后才变得明显。通过先转换为HTML，可以更准确地识别和移除这些潜在危险的脚本标签。
- **保持一致性**：在移除了`<script>`标签之后，再将HTML转回Markdown，确保最终存储的内容仍然是Markdown格式。这样做可以保持数据的一致性和可读性，方便后续处理或展示。
