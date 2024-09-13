package models

import "gorm.io/gorm"

// CommentInfo 评论信息
type CommentInfo struct {
	gorm.Model
	SubComments       []*CommentInfo `json:"sub_comments" gorm:"foreignKey:ParentCommentID"`   // 子评论
	ParentCommentInfo *CommentInfo   `json:"parent_comment" gorm:"foreignKey:ParentCommentID"` // 父评论
	ParentCommentID   *uint          `json:"parent_comment_id" gorm:"index"`                   // 父评论ID
	Content           string         `json:"content" gorm:"type:varchar(255)"`                 // 评论内容
	DiggCount         int            `json:"digg_count" gorm:"type:int"`                       // 点赞数
	CommentCount      int            `json:"comment_count" gorm:"type:int"`                    // 评论数
	Article           ArticleInfo    `json:"article" gorm:"foreignKey:ArticleID"`              // 文章
	ArticleID         uint           `json:"article_id" gorm:"index"`                          // 文章ID
	User              UserInfo       `json:"user" gorm:"foreignKey:UserID"`                    // 用户
	UserID            uint           `json:"user_id" gorm:"index"`                             // 用户ID
}
