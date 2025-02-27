package es_service

import (
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Fields []string
	Tag    string //标签搜索
}

// 用于排序，将接受到的排序字符串转化为es的排序参数
type SortField struct {
	Field     string
	Ascending bool
}

// GetForm 获取页码和每页显示的数量
// 生效于原值，需要使用指针
func (o *Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit
}
