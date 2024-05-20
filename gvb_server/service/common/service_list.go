package common

import (
	"gvb_server/global"
	"gvb_server/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool //是否开启debug模式,是否查看日志
}
// ComList 通用列表分页查询服务
func ComList[T any](model T , option Option) (list []T, count int64, err error) {
	
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	count = DB.Select("id").Find(&list).RowsAffected
	//偏移量
	offset := (option.Page - 1) * option.Limit
	//如果偏移量小于0，则从0开始
	if offset < 0 {
		offset = 0
	}
	//如果limit为0，则查询所有
	if option.Limit == 0 {
		option.Limit = -1
	}
	err = DB.Limit(option.Limit).Offset(offset).Find(&list).Error

	return list, count, err
}
