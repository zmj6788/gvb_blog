package cron_service

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_service"

	"gorm.io/gorm"
)

// 同步评论的点赞数据到mysql
// 总体思路：拿到所有的点赞数据，然后根据点赞的id去更新数据
func SyncCommentData() {
	commentDiggInfo := redis_service.NewCommentDigg().GetInfo()
	for key, count := range commentDiggInfo {
		var comment models.CommentModel
		err := global.DB.Take(&comment, key).Error
		if err != nil {
			global.Log.Error("评论不存在", err)
			continue
		}
		err = global.DB.Model(&comment).
		Update("digg_count", gorm.Expr("digg_count + ?", count)).Error
		if err != nil {
			global.Log.Error("更新点赞数失败", err)
			continue
		}
		global.Log.Infof("%s 更新点赞数成功, 新的点赞数为 %d", comment.Content, comment.DiggCount + count)
	}
	redis_service.NewCommentDigg().Clear()	
}
