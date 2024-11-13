package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/service/redis_service"
)

func main() {
	// 配置信息读取
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()
	global.Redis = core.ConnectRedis()
	global.ESClient = core.EsConnect()

	err := redis_service.Look("2DtL-5IBBOEDMw_pwTmW")
	if err != nil {
		global.Log.Error(err.Error())
	}
	global.Log.Info("浏览成功")

	global.Log.Info(redis_service.GetLook("1ztL-5IBBOEDMw_pRTnk"))
	global.Log.Info(redis_service.GetLookInfo())
	// redis_service.LookClear()

}
