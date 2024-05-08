package models

import "time"

// 自定义MODEL，没有用gorm的MODEL,因为我们不需要逻辑删除
type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"-"`                    // 更新时间
}
