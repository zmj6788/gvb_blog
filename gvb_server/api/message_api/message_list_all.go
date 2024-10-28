package message_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"

	"github.com/gin-gonic/gin"
)

// MessageListAllView 管理员查看所有消息列表
// @Tags 消息管理
// @Summary 管理员查看所有消息列表
// @Description 管理员查看所有消息列表
// @Param data query models.PageInfo    false  "查询参数"
// @Param token header string  true  "token"
// @Router /api/messages_all [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MessageModel]}
func (MessageApi) MessageListAllView(c *gin.Context) {

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//使用封装的列表分页查询服务
	list, count, err := common.ComList(
		models.MessageModel{},
		common.Option{
			PageInfo: page,
			Debug:    true,
		})
	if err != nil {
		res.FailWithMessage("消息列表获取失败", c)
		return
	}
	
	res.OkWithList(list, count, c)
}
