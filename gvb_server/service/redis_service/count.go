package redis_service

import (
	"gvb_server/global"
	"strconv"
)

// 优化 redis 存储数据，将浏览评论点赞方法合为一体
type CountDB struct {
	Index string // 索引
}

// Set 设置某一个数据，重复执行，重复累加
func (c CountDB) Set(id string) error {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	num++
	err := global.Redis.HSet(c.Index, id, num).Err()
	return err
}

// Get 获取某个的数据
func (c CountDB) Get(id string) int {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	return num
}

// GetInfo 取出数据
func (c CountDB) GetInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
