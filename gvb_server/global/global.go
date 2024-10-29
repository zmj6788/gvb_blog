package global

import (
	"gvb_server/config"

	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 配置信息存储全局变量，便于全局使用
var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
	Redis    *redis.Client
	ESClient *elastic.Client
)
