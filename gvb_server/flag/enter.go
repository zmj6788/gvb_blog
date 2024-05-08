package flag

import (
	sys_flag "flag"

	"github.com/sirupsen/logrus"
)

type Option struct {
	DB bool
}

//Parse 解析命令行参数
func Parse() Option {
	//为db设置默认值，默认不进行表结构迁移
	db := sys_flag.Bool("db", false, "初始化数据库")
	//解析命令行参数写入注册的db中
	sys_flag.Parse()
	logrus.Info("db:", *db)
	return Option{
		DB: *db,
	}
}

// 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

// 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
}
