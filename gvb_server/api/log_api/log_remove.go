package log_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"

	"github.com/gin-gonic/gin"
)
// LogRemoveListView 批量删除日志
// @Tags 日志管理
// @Summary 批量删除日志
// @Description 批量删除日志
// @Param data body models.RemoveRequest    true  "图片id列表"
// @Router /api/logs [delete]
// @Produce json
// @Success 200 {object} res.Response{}
//图片删除接口，接受一个删除请求，请求参数为idlist的json数据
func (LogApi) LogRemoveListView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var list []log_stash.LogStashModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("日志不存在", c)
		return
	}
	global.DB.Delete(&list)
	res.OkWithMessage(fmt.Sprintf("共删除 %d 个日志", count), c)

}
