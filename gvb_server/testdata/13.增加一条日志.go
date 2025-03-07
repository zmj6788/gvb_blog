package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/plugins/log_stash"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.Initgorm()

	log := log_stash.New("127.0.0.1", "qwqwqw")
	log.Info("日志创建测试")
}
