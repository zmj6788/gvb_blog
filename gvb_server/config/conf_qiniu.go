package config
// 七牛云
type QiNiu struct {
	Enable    bool    `json:"enable" yaml:"enable"` // 是否启用七牛云来缓存图片
	AccessKey string  `json:"access_key" yaml:"access_key"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"` // 存储空间
	CDN       string  `json:"cdn" yaml:"cdn"`       // 访问图片的地址的前缀
	Zone      string  `json:"zone" yaml:"zone"`     // 存储的地区
	Size      float64 `json:"size" yaml:"size"`     // 存储的大小限制，单位为MB
}
