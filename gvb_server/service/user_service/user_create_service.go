package user_service

import (
	"errors"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/untils"
	"gvb_server/untils/pwd"
)

// CreateUser 创建用户
func (UserService) CreateUser(nickName string, userName string, password string, role ctype.Role, email string, ip string) error{
	// 判断用户名是否重复
	var user models.UserModel
	err := global.DB.Take(&user,"user_name = ?",userName).Error
	if err == nil {
		// 用户名重复
		return errors.New("用户名已存在")
	}
	hashPwd := pwd.HashPwd(password)
	avatar := "http://localhost:8080/api/uploads/avatar/1.jpg"
	addr := untils.GetAddr(ip)
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         ip,
		Addr:       addr,
		SignStatus: ctype.SignEmail,

	}).Error
	if err != nil {
		return err
	}
	return nil
}
