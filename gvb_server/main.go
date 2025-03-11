package main

import (
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
	"gvb_server/service/cron_service"
	"gvb_server/untils"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server API文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	// 配置信息读取
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()
	//数据库连接
	global.DB = core.Initgorm()
	//redis连接
	global.Redis = core.ConnectRedis()
	//es连接
	global.ESClient = core.EsConnect()
	//连接ip地址数据库
	core.InitAddrDB()
	defer global.AddrDB.Close()
	//命令行参数绑定迁移表结构函数
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		//控制迁移表结构后退出
		return
	}
	// 定时任务初始化，同步文章数据以及评论数据从redis到es或mysql中
	cron_service.CronInit()
	//路由初始化
	router := routers.InitRouter()
	//启动服务
	addr := global.Config.System.Addr()
	untils.PrintSystem()
	err  := router.Run(addr)
	if err != nil {
	}
	router.Run(addr)
}