package config

import "fmt"
//qq
type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` // 登陆后的回调地址
}
// 获取QQ登陆地址
func (q QQ) GetPath() string {
	if q.AppID == "" || q.Key == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=#{q.AppID}&redirect_url=#{q.Redirect}")
}
