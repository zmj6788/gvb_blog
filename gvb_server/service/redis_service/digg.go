package redis_service

import (
	"gvb_server/global"
	"strconv"
)

const diggPrefix = "digg"

/*
数据结构

例子：HSET myhash field1 "value1", HSET myhash field2 "value2"
结构：key -> {field1: value1, field2: value2, ...}

*/
// Digg 点赞一篇文章
func Digg(id string) error {
	// 获取点赞数
	// 第一次若没有获取到num被赋值为0
	num, _ := global.Redis.HGet(diggPrefix, id).Int()
	num++
	// 更新点赞数
	err := global.Redis.HSet(diggPrefix, id, num).Err()
	return err
}

// GetDigg 获取某一篇文章下的点赞数
func GetDigg(id string) int {
	num, _ := global.Redis.HGet(diggPrefix, id).Int()
	return num
}

// GetDiggInfo 取出点赞数据
func GetDiggInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(diggPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

// DiggClear 清空点赞数据
func DiggClear() {
	global.Redis.Del(diggPrefix)
}
