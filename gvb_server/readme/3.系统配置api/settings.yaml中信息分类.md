
settings.yaml

settings.yaml中编写了许多信息，他们的共同点是都是配置信息

类如：mysql，system，logger的配置信息使用在后端项目启动时的操作

类如:   site_info，email，jwt，qi_niu，qq的配置信息适用于，后端项目启动后，客户端后台管理系

统中管理员使用这些配置信息对客户端进行统一的设置以及管理

```
mysql:

  host: 127.0.0.1

  port: 3306

  config: ""

  db: gvb_db

  user: root

  password: "123456"

  log_level: dev

system:

  host: 0.0.0.0

  port: 8080

  env: release

logger:

  level: info

  prefix: '[gvb]'

  director: log

  show_line: true

  log_in_console: true

site_info:

  created_at: "2024-04-23"

  bei_an: 暂无备案

  title: 张明杰的个人博客

  qq_image: https://image.baidu.com/search/albumsdetail?tn=albumsdetail&word=%E5%AE%A0%E7%89%A9%E5%9B%BE%E7%89%87&fr=albumslist&album_tab=%E5%8A%A8%E7%89%A9&album_id=688&rn=30

  version: v1.0.0

  email: 2124427385@qq.com

  wechat_image: https://image.baidu.com/search/albumsdetail?tn=albumsdetail&word=%E6%B8%90%E5%8F%98%E9%A3%8E%E6%A0%BC%E6%8F%92%E7%94%BB&fr=albumslist&album_tab=%E8%AE%BE%E8%AE%A1%E7%B4%A0%E6%9D%90&album_id=409&rn=30

  name: zmj

  job: 后端开发

  addr: 河南郑州

  slogan: mj_blog

  web: https://www.mj.com

  bilibili_url: https://space.bilibili.com/3546656056281480?spm_id_from=333.788.0.0

  gitee_url: https://gitee.com/zhang-mingjie970853

  github_url: https://github.com/zmj6788

email:

  host: smtp.qq.com

  port: 465

  user: xxx@qq.com

  password: xxx

  default_from_email: 枫枫知道

  use_ssl: xx

  use_tls: xx

jwt:

  secret: xxx

  expires: 48

  issuer: xx

qi_niu:

  access_key: UMM-xx-FdRIf-xx

  secret_key: xx

  bucket: fengfengzhidao

  cdn: xx://XX.XX.com/

  zone: xx

  size: 5

qq:

  app_id: "xx"

  key: xx

  redirect: http://xx.xx.com/login?flag=qq
```