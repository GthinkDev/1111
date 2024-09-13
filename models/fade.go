package models

import "gorm.io/gorm"

// FadeBack 用户反馈表
type FadeBack struct {
	gorm.Model
	Content      string `json:"content"`       // 内容
	Email        string `json:"email"`         //
	ApplyContent string `json:"apply_content"` // 回复的内容
	IsApply      bool   `json:"is_apply"`      // 是否回复
}
