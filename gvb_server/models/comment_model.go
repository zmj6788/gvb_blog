package models

// 评论表
type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  //子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` //父评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                               //父评论ID
	Conent             string          `gorm:"size:256" json:"conent"`                          //评论内容
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`              //点赞数
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`           //子评论数
	Article            ArticleModel    `gorm:"foreignKey:ArticleID" json:"-"`                   //关联的文章
	ArticleID          uint            `json:"article_id"`                                      //文章id
	User               UserModel       `json:"user"`                                            //关联的用户
<<<<<<< HEAD
	UserID             uint            `json:"user_id"`                          //评论的用户ID
=======
	UserID             uint            `gorm:"size:10" json:"user_id"`                          //评论的用户ID
>>>>>>> 2f9e4d1a6a0ab0002a002517dace0301441cd6ca
}
