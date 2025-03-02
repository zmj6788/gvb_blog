package comment_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_service"
	"gvb_server/service/redis_service"
	"gvb_server/untils/jwts"

	"github.com/gin-gonic/gin"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 父评论id
}

// CommentCreateView 发布评论
// @Tags 评论管理
// @Summary 发布评论
// @Description 发布评论
// @Param data body CommentRequest    true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/comments [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (CommentApi) CommentCreateView(c *gin.Context) {
	// 获取评论数据
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}
	// 获取用户id
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	// 文章是否存在
	_, err = es_service.CommDetail(cr.ArticleID)
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}
	// 父评论是否存在
	if cr.ParentCommentID != nil {
		var ParentComment models.CommentModel
		err = global.DB.Take(&ParentComment, "id = ?", cr.ParentCommentID).Error
		if err != nil {
			res.FailWithMessage("父评论不存在", c)
			return
		}
		if ParentComment.ArticleID != cr.ArticleID {
			res.FailWithMessage("父评论不属于该文章", c)
			return
		}
		// 父评论下的评论数量+1
		global.DB.Model(&ParentComment).UpdateColumn("comment_count", ParentComment.CommentCount+1)
	}
	// 没有父评论，评论成功，评论数据入库

	err = global.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	}).Error

	if err != nil {
		res.FailWithMessage("评论失败", c)
		return
	}
	// 文章下的评论数量+1
	redis_service.Comment(cr.ArticleID)
	res.OkWithMessage("评论成功", c)
}
