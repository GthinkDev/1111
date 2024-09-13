package models

import (
	"gorm.io/gorm"
)

// ImageInfo 图片模型
type ImageInfo struct {
	gorm.Model
	Path string `json:"path"`                          // 图片的路径
	Hash string `json:"hash"`                          // 图片的hash值
	Name string `gorm:"type:varchar(100)" json:"name"` // 图片的名称
}
