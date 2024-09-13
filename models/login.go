package models

import "gorm.io/gorm"

type LoginData struct {
	gorm.Model
	UserID   string   `json:"user_id"`
	UserInfo UserInfo `json:"-" gorm:"foreignKey:UserID"`
	IP       string   `json:"ip" gorm:"size:32"`
	NickName string   `json:"nick_name" gorm:"size:64"`
	Token    string   `json:"token" gorm:"size:256"`
	Device   string   `json:"device" gorm:"size:256"`
	Addr     string   `json:"addr" gorm:"size:256"`
}
