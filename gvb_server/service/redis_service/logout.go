package redis_service

import (
	"gvb_server/global"
	"gvb_server/untils"
	"time"
)


/*
数据结构

例子：SET mykey "Hello, Redis!"
结构：key -> value

*/
const prefix = "logout_"
func Logout(token string, diff time.Duration) error {
	//将注销用户的token放入redis中

	err := global.Redis.Set(prefix+token, "", diff).Err()
	return err
}

func CheckLogout(token string) bool {
	//验证token是否在redis注销列表token中
	keys := global.Redis.Keys(prefix + "*").Val()
	// global.Log.Info(keys)
	if untils.InList(prefix+token, keys) {
		return true
	}
	return false
}
