package log_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"
	"gvb_server/service/common"

	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	models.PageInfo
	Level log_stash.Level `form:"level"`
}
// LogListView 日志列表
// @Tags 日志管理
// @Summary 日志列表
// @Description 日志列表
// @Param data query LogRequest    false  "查询参数"
// @Router /api/logs [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[log_stash.LogStashModel]}
func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	c.ShouldBindQuery(&cr)
	list, count, _ := common.ComList(log_stash.LogStashModel{Level: cr.Level}, common.Option{
		PageInfo: cr.PageInfo,
		Debug:    true,
		Likes:    []string{"ip", "addr"},
	})
	res.OkWithList(list, count, c)
}
