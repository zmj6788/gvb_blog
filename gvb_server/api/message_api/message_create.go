package message_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type MessageCreateRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` //发送者id
	RevUserID  uint   `json:"rev_user_id" binding:"required"`  //接收者id
	Content    string `json:"content"  binding:"required"`                        //消息内容
}

// MessageCreateView 发布消息
// @Tags 消息管理
// @Summary 发布消息
// @Description 发布消息
// @Param data body MessageCreateRequest    true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/messages [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (MessageApi) MessageCreateView(c *gin.Context) {
	//发送者id应该是从token解析出当前登录的用户
	//为了方标测试时不用一直切换token，使用当前方式
	//可以在前端获得当前用户id搭配当前模式，测试方便
	var req MessageCreateRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.FailWithError(err, &req, c)
		return
	}

	//获取到发布消息请求后
	//查看发送者和接收者是否存在
	var sendUser, revUser models.UserModel
	err = global.DB.Take(&sendUser, req.SendUserID).Error
	if err != nil {
		res.FailWithMessage("发送者不存在", c)
		return
	}
	err = global.DB.Take(&revUser, req.RevUserID).Error
	if err != nil {
		res.FailWithMessage("接受者不存在", c)
		return
	}

	//发送者接收者都存在，消息入库
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       req.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        req.RevUserID,
		RevUserNickName:  revUser.NickName,
		RevUserAvatar:    revUser.Avatar,
		IsRead:           false,
		Content:          req.Content,
	}).Error
		
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("消息发送失败", c)
		return
	}
	
	res.OkWithMessage("发布消息成功", c)
}
