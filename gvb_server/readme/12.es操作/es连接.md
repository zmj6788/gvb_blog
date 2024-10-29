

```
func EsConnect() *elastic.Client {

  var err error

  sniffOpt := elastic.SetSniff(false)

  host := "http://172.17.97.252:9200"

  c, err := elastic.NewClient(

    elastic.SetURL(host),

    sniffOpt,

    elastic.SetBasicAuth("elastic", "PDgvGi30SU0=sesPOHCB"),

  )

  if err != nil {

    logrus.Fatalf("es连接失败 %s", err.Error())

  }

  logrus.Info("es连接成功")

  return c

}
```