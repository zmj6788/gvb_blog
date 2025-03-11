package cron_service

import (
	"time"

	"github.com/robfig/cron/v3"
)

func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	// 每日0点执行
	Cron.AddFunc("0 0 0 * * *", SyncArticleData)
	Cron.AddFunc("0 0 0 * * *", SyncCommentData)
	Cron.Start()
}
