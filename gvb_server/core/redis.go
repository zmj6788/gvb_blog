package core

import (
	"context"
	"gvb_server/global"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

// 默认使用 db=0 ，连接redis
func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

// 指定db连接redis
func ConnectRedisDB(db int) *redis.Client {
	redisConfig := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr(),
		Password: redisConfig.Password, // no password set
		DB:       db,                   // use default DB
		PoolSize: redisConfig.PoolSize, // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Errorf("redis连接失败: %s", redisConfig.Addr())
		return nil
	}
	logrus.Info("redis连接成功")
	return rdb
}
