package chat_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

// ChatListView 群聊消息列表
// @Tags 群聊管理
// @Summary 群聊消息列表
// @Description 群聊消息列表
// @Router /api/chat_groups_records [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (ChatApi) ChatListView(c *gin.Context) {

	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 排序,先获取最新的聊天记录
	page.Sort = "created_at desc"
	//使用封装的列表分页查询服务
	list, count, err := common.ComList(
		models.ChatModel{},
		common.Option{
			PageInfo: page,
		})
	if err != nil {
		res.FailWithMessage("群聊消息列表获取失败", c)
		return
	}
	// 排除 list字段,json-filter空值问题解决
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ChatModel, 0)
		res.OkWithList(list, count, c)
		return
	}

	res.OkWithList(data, count, c)
}
