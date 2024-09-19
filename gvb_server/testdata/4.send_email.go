package main

import (
	"gvb_server/core"
	"gvb_server/global"

	"gopkg.in/gomail.v2"
)

type Subject string

//预定义的邮件主题
const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alarm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}
func NewNote() Api {
	return Api{
		Subject: Note,
	}
}
func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

// send 邮件发送  发给谁，主题，正文
func send(name, subject, body string) error {
	e := global.Config.Email
	return sendMail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail,
		subject,
		body,
	)
}

func sendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName)) // 发件人邮箱，发件人名字
	m.SetHeader("To", mailTo)                                // 发送给谁
	m.SetHeader("Subject", subject)                          // 主题
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, authCode)
	err := d.DialAndSend(m)
	return err
}

func main() {
	// 配置信息读取
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()
	// 发送邮件
	NewCode().Send("2124427385@qq.com", "<p>122788</p>")
}
