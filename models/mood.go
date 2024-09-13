package models

import (
	"gin-blog/models/ctype"
	"gorm.io/gorm"
)

// MoodInfo 心情表
type MoodInfo struct {
	gorm.Model
	UserID  uint        `json:"user_id"`                    // 用户ID
	User    UserInfo    `json:"-" gorm:"foreignKey:UserID"` // 用户信息
	Avatar  string      `json:"avatar"`                     // 头像
	IP      string      `json:"ip" gorm:"size:32"`          // IP地址
	Addr    string      `json:"addr" gorm:"size:256"`       // 地址
	Content string      `json:"content" gorm:"size:256"`    // 心情内容
	Drawing ctype.Array `json:"drawing" gorm:"type:longtext"`
}
