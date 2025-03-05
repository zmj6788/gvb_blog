package comment_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redis_service"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `form:"article_id"`
}

// CommentListView 评论列表
// @Tags 评论管理
// @Summary 评论列表
// @Description 评论列表
// @Param data query CommentListRequest    false  "查询参数"
// @Router /api/comments [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.CommentModel]}
func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)
	res.OkWithData(filter.Select("c", rootCommentList), c)
}

func FindArticleCommentList(articleID string) (RootCommentList []*models.CommentModel) {
	// 先把文章下的根评论查出来
	global.DB.Preload("User").Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	// 获取评论点赞数据
	diggInfo := redis_service.NewCommentDigg().GetInfo()
	// 遍历根评论，递归查根评论下的所有子评论
	for _, model := range RootCommentList {
		var subCommentList, newSubCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		for _, commentModel := range subCommentList {
			digg := diggInfo[fmt.Sprintf("%d", commentModel.ID)]
			commentModel.DiggCount = commentModel.DiggCount + digg
			newSubCommentList = append(newSubCommentList, commentModel)
		}
		// 获取根评论点赞数据
		modelDigg := diggInfo[fmt.Sprintf("%d", model.ID)]
		// 根评论点赞数据赋值
		model.DiggCount = model.DiggCount + modelDigg
		model.SubComments = newSubCommentList
	}
	return
}

// FindSubComment 递归查评论下的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
	return
}
// FindSubCommentCount 查评论下的子评论个数
func FindSubCommentCount(model models.CommentModel) (subCommentList []models.CommentModel){
	FindSubComment(model, &subCommentList)
	return subCommentList
}
