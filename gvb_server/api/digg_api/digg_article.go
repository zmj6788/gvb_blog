package digg_api

import (
	"gvb_server/models/res"
	"gvb_server/service/redis_service"

	"github.com/gin-gonic/gin"
)

type DiggRequest struct {
	ID string `uri:"id" json:"id" form:"id"`
}
// DiggArticleView 文章点赞
// @Tags 文章管理
// @Summary 文章点赞
// @Description 文章点赞
// @Param data body DiggRequest    true  "表示多个参数"
// @Router /api/digg/article [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (DiggApi) DiggArticleView(c *gin.Context) {
	var req DiggRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}
	// 可以校验id是否存在或长度是否符合
	err = redis_service.Digg(req.ID)
	if err != nil {
		res.FailWithMessage("点赞失败", c)
		return
	}
	res.OkWithMessage("点赞成功", c)
}