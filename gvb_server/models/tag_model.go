package models

// 标签表
type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`            // 标签名
	Articles []ArticleModel `gorm:"many2many:article_tag_models;" json:"-"` // 标签文章
}
