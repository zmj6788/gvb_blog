package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/untils/jwts"
)

func main() {
	// 配置信息读取
	core.InitConf()
	//日志初始化
	global.Log = core.InitLogger()
	// 生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{NickName: "所悟皆是空", UserID: 1, Username: "张明杰", Role: 1})
	fmt.Println(token, err)

	//解析token
	customClaims, err := jwts.ParseToken(token)
	fmt.Println(customClaims, err)

}
