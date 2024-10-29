### 索引 (Index)

- **定义**：索引是Elasticsearch中用于存储文档的地方。它类似于关系数据库中的表
相当于表明
### 映射 (Mapping)

- **定义**：映射描述了索引中文档的字段及其对应的类型和属性
相当于规定表内类型的属性

### 创建索引

通过Index和Mapping创建索引

### 删除索引

通过index删除索引

### 创建一条数据

实例化对象通过索引index创建

### 查询数据

通过索引index分页查询
通过索引index分页查询特定字段
限制字段使用
```
FetchSourceContext(elastic.NewFetchSourceContext(true).Include("title")).
```
使用
```
Source(`{"_source": ["title"]}`).
和 
Source(&elastic.SourceConfig{Includes: []string{"title"}}). 
均无法使用
```

### 更新数据

通过数据id查找对应数据更新
- "_id": "GOQg15IBSFVhxiV54o14",
### 删除数据

通过数据id删除数据
```
count, err := Remove([]string{"GOQg15IBSFVhxiV54o14", "IuRK15IBSFVhxiV5OY1J", "GeQq15IBSFVhxiV5R418", "GuQr15IBSFVhxiV5kY3W"})

  fmt.Println(count, err)
```

