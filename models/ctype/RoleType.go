package ctype

import "encoding/json"

type Role int

const (
	Admin   Role = 1 // Admin 管理员
	User    Role = 2 // User 普通用户
	Guest   Role = 3 // Guest 游客
	Disable Role = 4 // Disable 禁用
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.string())
}

func (r Role) string() string {
	var str string
	switch r {
	case Admin:
		str = "管理员"
	case User:
		str = "用户"
	case Guest:
		str = "游客"
	case Disable:
		str = "被禁用"
	default:
		str = "未知"
	}
	return str
}
