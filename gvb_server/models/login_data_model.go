package models

//统计用户登录数据
type LoginDataModel struct {
	MODEL
	UserID    uint      `json:"user_id"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"` // ip
	NickName  string    `gorm:"size:42" json:"nick_name"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` // 设备
	Addr      string    `gorm:"size:64" json:"addr"`
}
