package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456", // no password set
		DB:       0,       // use default DB
		PoolSize: 100,     // 连接池大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Error(err)
		return
	}

}
func main() {
	err := rdb.Set("xxx1", "value1", 10*time.Second).Err()
	fmt.Println(err)
	cmd := rdb.Keys("*")
	keys, err := cmd.Result()
	fmt.Println(keys, err)
}
