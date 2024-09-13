package models

import "time"

// Collects 收藏表
type Collects struct {
	UserID    uint        `gorm:"primaryKey"`
	UserInfo  UserInfo    `gorm:"foreignKey:UserID"`    // 外键关联用户信息
	ArticleID uint        `gorm:"primaryKey"`           // 外键关联文章ID
	Article   ArticleInfo `gorm:"foreignKey:ArticleID"` // 外键关联文章信息
	CreatedAt time.Time
}
