package chat_api

import (
	"encoding/json"
	"fmt"
	"gvb_server/models/res"
	"net/http"
	"strings"
	"time"

	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

var ConnGroupMap = map[string]ChatUser{}

type MsgType int

const (
	TextMsg    MsgType = 1
	ImageMsg   MsgType = 2
	SystemMsg  MsgType = 3
	InRoomMsg  MsgType = 4
	OutRoomMsg MsgType = 5
)

type GroupRequest struct {
	Content string  `json:"content"`  // 聊天的内容
	MsgType MsgType `json:"msg_type"` // 聊天类型
}
type GroupResponse struct {
	NickName string    `json:"nick_name"` // 随机生成
	Avatar   string    `json:"avatar"`    // 头像
	MsgType  MsgType   `json:"msg_type"`  // 聊天类型
	Content  string    `json:"content"`   // 聊天的内容
	Date     time.Time `json:"date"`      // 消息的时间
}

func (ChatApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 鉴权 true表示放行，false表示拦截
			return true
		},
	}
	// 将http升级至websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	addr := conn.RemoteAddr().String()
	// 随机生成昵称
	nickName := randomname.GenerateName()
	// 根据昵称首字头像关联
	nickNameFirst := string([]rune(nickName)[0])
	avatar := fmt.Sprintf("uploads/chat_avatar/%s.png", nickNameFirst)
	// 连接成功，将用户信息保存到map中，群聊列表中，关联地址
	chatUser := ChatUser{
		Conn:     conn,
		NickName: nickName,
		Avatar:   avatar,
	}
	ConnGroupMap[addr] = chatUser

	logrus.Infof("%s 连接成功", addr)
	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天
			SendGroupMsg(GroupResponse{
				Content: fmt.Sprintf("%s 离开聊天室", chatUser.NickName),
				Date:    time.Now(),
			})
			break
		}
		// 进行参数绑定
		var request GroupRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			// 参数绑定失败
			continue
		}
		// 判断类型
		switch request.MsgType {
		case TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				continue
			}
			SendGroupMsg(GroupResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				Content:  request.Content,
				MsgType:  TextMsg,
				Date:     time.Now(),
			})
		case InRoomMsg:
			SendGroupMsg(GroupResponse{
				Content: fmt.Sprintf("%s 进入聊天室", chatUser.NickName),
				Date:    time.Now(),
			})
		}

	}
	defer conn.Close()
	// 断开连接，删除map中对应用户数据
	delete(ConnGroupMap, addr)
}

// SendGroupMsg 群聊功能
func SendGroupMsg(response GroupResponse) {
	byteData, _ := json.Marshal(response)
	for _, chatUser := range ConnGroupMap {
		chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}
