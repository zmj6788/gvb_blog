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

	digg := redis_service.NewDigg()
	digg.Set("FhPOWZUBduh9nTRpg7BQ")
	
	global.Log.Info("点赞成功")

	// global.Log.Info(redis_service.GetDigg("1ztL-5IBBOEDMw_pRTnk"))
	global.Log.Info(digg.GetInfo())
	// redis_service.DiggClear()

}
