
```
// 创建一个桶
bulk := global.ESClient.Bulk()
for _, indexData := range indexList {
  req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
  bulk.Add(req)
}
result, err := bulk.Do(context.Background())
```