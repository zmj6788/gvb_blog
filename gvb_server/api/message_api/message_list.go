package message_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/untils/jwts"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	SendUserID       uint      `json:"send_user_id"` // 发送人id
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"` // 接收人id
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`       // 消息内容
	CreatedAt        time.Time `json:"created_at"`    // 最新的消息时间
	MessageCount     int       `json:"message_count"` // 消息条数
}

type MessageGroup map[uint]*Message

// MessageListView 用户消息列表
// @Tags 消息管理
// @Summary 用户消息列表
// @Description 用户消息列表
// @Param token header string  true  "token"
// @Router /api/messages [get]	
// @Produce json
// @Success 200 {object} res.Response{data=[]Message}
func (MessageApi) MessageListView(c *gin.Context) {
	// 获取当前登录用户ID
	//通过中间件拿到解析token后的用户信息
	_claims, _ := c.Get("claims")
	//将获取的值转换为 jwts.CustomClaims 类型，并存储在变量 claims 中
	claims := _claims.(*jwts.CustomClaims)

	//查询到用户相关消息
	var messageGroup = MessageGroup{}
	var messageList []models.MessageModel
	var messages []Message
	global.DB.Order("created_at asc").
		Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, model := range messageList {
		//我们可以知道当前信息列表内数据，发送方或接受方一定有一个id等于当前登录用户id
		//那么两个id和相同的消息数据一定是互为聊天对方
		// 1 2  2 1
    // 1 3  3 1 是一组
		message := Message{
			SendUserID:       model.SendUserID,
			SendUserNickName: model.SendUserNickName,
			SendUserAvatar:   model.SendUserAvatar,
			RevUserID:        model.RevUserID,
			RevUserNickName:  model.RevUserNickName,
			RevUserAvatar:    model.RevUserAvatar,
			Content:          model.Content,          //最新消息内容
			CreatedAt:        model.CreatedAt,        //最新消息时间
			MessageCount:     1,                      //双方消息总数
		}
		// 根据发送用户ID和接收用户ID生成一个唯一的idNum，并检查messageGroup映射中是否存在该键。如果不存在，
		// 则将新的message指针添加到messageGroup中；如果存在，则跳过当前循环迭代。
		idNum := model.SendUserID + model.RevUserID
		val, ok := messageGroup[idNum]
		if !ok {
			// 不存在
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	//我们最终返回的数据，是双方最新的一条消息
	//所有与我有关的最新消息的集合
	res.OkWithData(messages, c)

}
