package main

import (
	"gvb_server/core"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	// 配置信息读取
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()
	//数据库连接
	global.DB = core.Initgorm()
	//命令行参数绑定迁移表结构函数
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		//控制迁移表结构后退出
		return
	}
	//路由初始化
	router := routers.InitRouter()
	//启动服务
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在: %s", addr)
	err  := router.Run(addr)
	if err != nil {
	}
	router.Run(addr)
}