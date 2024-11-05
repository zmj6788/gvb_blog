package main

import (
	"gvb_server/core"
	"gvb_server/flag"
	"gvb_server/global"
)

func init() {
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()
	//数据库连接
	global.DB = core.Initgorm()
	//redis连接
	global.Redis = core.ConnectRedis()
	//es连接
	global.ESClient = core.EsConnect()
	//命令行参数绑定迁移表结构函数
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		//控制迁移表结构后退出
		return
	}
}

func main() {

}
