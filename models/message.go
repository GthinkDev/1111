package models

import "gorm.io/gorm"

// MessageInfo 消息表
type MessageInfo struct {
	gorm.Model
	SendUserID       uint     `json:"send_user_id" gorm:"primaryKey"` // 发送者ID
	SendUserInfo     UserInfo `json:"-" gorm:"foreignKey:SendUserID"` // 发送者信息
	SendUserNickName string   `json:"send_user_nick_name" gorm:"64"`  // 发送者昵称
	SendUserAvatar   string   `json:"send_user_avatar"`               // 发送者头像

	RecvUserID       uint     `json:"recv_user_id" gorm:"primaryKey"` // 接收者ID
	RecvUserInfo     UserInfo `json:"-" gorm:"foreignKey:RecvUserID"` // 接收者信息
	RecvUserNickName string   `json:"recv_user_nick_name" gorm:"64"`  // 接收者昵称
	RecvUserAvatar   string   `json:"recv_user_avatar"`               // 接收者头像
	IsRead           bool     `json:"is_read" gorm:"default:false"`   // 是否已读
	Content          string   `json:"content"`                        // 消息内容
}
