package ctype

import "encoding/json"

type SignStatus int

const (
	QQ     SignStatus = 1 // QQ qq
	WeChat SignStatus = 2 // WeChat 微信
	Email  SignStatus = 3 // Email 邮箱
	Phone  SignStatus = 4 // Phone 手机号码
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.string())
}

func (s SignStatus) string() string {
	var str string
	switch s {
	case QQ:
		str = "qq"
	case WeChat:
		str = "微信"
	case Email:
		str = "邮箱"
	case Phone:
		str = "手机号码"
	default:
		str = "未知"
	}
	return str
}
