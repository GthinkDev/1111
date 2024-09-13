package models

import "gorm.io/gorm"

// TagInfo 标签
type TagInfo struct {
	gorm.Model
	Title    string        `json:"title" gorm:"size:16"`           // 标签的名称
	Articles []ArticleInfo `json:"-" gorm:"many2many:article_tag"` // 标签关联的文章
}
