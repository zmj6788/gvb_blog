package models

// 标签表
type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`            // 标签名
}
