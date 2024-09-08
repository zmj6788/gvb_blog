package flag

import (
	sys_flag "flag"
)

type Option struct {
	DB   bool
	User string // 创建用户 -u user -u admin
}

// Parse 解析命令行参数
func Parse() Option {
	//为db设置默认值，默认不进行表结构迁移
	db := sys_flag.Bool("db", false, "初始化数据库")
	user := sys_flag.String("u", "", "创建用户")
	//解析命令行参数写入注册的db中
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return option.DB
}

// 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "user" || option.User == "admin" {
		CreateUser(option.User)
		return
	}
	if option.User != "" {
		sys_flag.Usage()
	}
}
