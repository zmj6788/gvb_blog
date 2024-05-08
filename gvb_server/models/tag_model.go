package models

// 标签表
type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`            // 标签名
<<<<<<< HEAD
	Articles []ArticleModel `gorm:"many2many:article_tag_models;" json:"-"` // 标签文章
=======
	Articles []ArticleModel `gorm:"many2many:article_tag;" json:"-"` // 标签文章
>>>>>>> 2f9e4d1a6a0ab0002a002517dace0301441cd6ca
}
