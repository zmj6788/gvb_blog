package models

// 广告表
type AdvertModel struct {
	MODEL
	Title  string `gorm:"size:32" json:"title"` //广告标题
	Href   string `json:"href"`                 //跳转链接
	Images string `json:"images"`               //广告图片
	IsShow bool   `json:"is_show"`              //是否显示
}
