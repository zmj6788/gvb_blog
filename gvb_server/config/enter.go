package config

// 存储配置信息结构体
type Config struct {
	MySQL    MySQL    `yaml:"mysql"`
	System   System   `yaml:"system"`
	Logger   Logger   `yaml:"logger"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	QQ       QQ       `yaml:"qq"`
	Upload  Upload  `yaml:"upload"`
}
