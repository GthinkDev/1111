package models

import "gorm.io/gorm"

type LogInfo struct {
	gorm.Model
	IP       string   `json:"ip" gorm:"size:32"`
	Addr     string   `json:"addr" gorm:"size:32"`
	UserID   uint     `json:"user_id"`
	UserInfo UserInfo `json:"-" gorm:"foreignKey:UserID"`
	NickName string   `json:"nick_name" gorm:"size:32"`
	Level    string   `json:"level" gorm:"size:32"`
	Message  string   `json:"message" gorm:"size:32"`
}
