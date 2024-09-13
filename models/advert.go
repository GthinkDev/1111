package models

import "gorm.io/gorm"

// AdvertInfo 广告
type AdvertInfo struct {
	gorm.Model
	Title  string `json:"title" gorm:"size:32"` // 标题
	Href   string `json:"href"`                 // 链接
	Images string `json:"images"`               // 图片
	IsShow bool   `json:"is_show"`              // 是否显示
}
