package config

type QiNiu struct {
	Enable    bool    `json:"enable" yaml:"enable"`         // 是否启用七牛云
	AccessKey string  `json:"access_key" yaml:"access_key"` // 七牛云AccessKey
	SecretKey string  `json:"secret_key" yaml:"secret_key"` // 七牛云SecretKey
	Bucket    string  `json:"bucket" yaml:"bucket"`         // 七牛云存储同
	CDN       string  `json:"cdn" yaml:"cdn"`               // 七牛云 访问图片地址的前缀
	Zone      string  `json:"zone" yaml:"zone"`             // 七牛云 存储的地区
	Size      float64 `json:"size" yaml:"size"`             // 七牛云文件存储的大小限制，单位 mb
}
