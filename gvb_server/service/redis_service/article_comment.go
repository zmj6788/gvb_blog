package redis_service

import (
	"gvb_server/global"
	"strconv"
)

const commentPrefix = "comment"

// Look 浏览某一篇文章
func Comment(id string) error {
	num, _ := global.Redis.HGet(commentPrefix, id).Int()
	num++
	err := global.Redis.HSet(commentPrefix, id, num).Err()
	return err
}

// GetLook 获取某一篇文章下的浏览数
func GetComment(id string) int {
	num, _ := global.Redis.HGet(commentPrefix, id).Int()
	return num
}

// GetLookInfo 取出浏览量数据
func GetCommentInfo() map[string]int {
	var CommentInfo = map[string]int{}
	maps := global.Redis.HGetAll(commentPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		CommentInfo[id] = num
	}
	return CommentInfo
}

func CommentClear() {
	global.Redis.Del(commentPrefix)
}
