package models

import (
	"gin-blog/models/ctype"

	"gorm.io/gorm"
)

// ArticleInfo 文章模型
type ArticleInfo struct {
	gorm.Model
	Title         string        `json:"title" gorm:"size:32"`                     // 文章标题
	Abstract      string        `json:"abstract"`                                 // 文章摘要
	Content       string        `json:"content"`                                  // 文章内容
	LookCount     int           `json:"look_count"`                               // 文章浏览量
	CommentCount  int           `json:"comment_count"`                            // 文章评论量
	DiggCount     int           `json:"digg_count"`                               // 文章点赞量
	CollectsCount int           `json:"collects_count"`                           // 文章收藏量
	TagInfo       []TagInfo     `json:"tag_info" gorm:"many2many:article_tag"`    // 文章标签
	CommentInfo   []CommentInfo `json:"comment_info" gorm:"foreignKey:ArticleID"` // 文章评论
	UserInfo      UserInfo      `json:"-" gorm:"foreignKey:UserID"`               // 文章作者
	UserID        uint          `json:"user_id"`                                  // 用户 id
	Category      string        `json:"category" gorm:"size:24"`                  // 文章分类
	Source        string        `json:"source"`                                   // 文章来源
	Link          string        `json:"link"`                                     // 文章链接
	Image         ImageInfo     `json:"-" gorm:"foreignKey:ImageID" `             // 文章封面
	ImageID       uint          `json:"image_id"`                                 // 文章封面ID
	ImagePath     string        `json:"image_path"`                               // 文章封面路径
	NickName      string        `json:"nick_name" gorm:"size:40"`                 // 文章作者昵称
	Tags          ctype.Array   `json:"tags" gorm:"type:string; size:64"`         // 文章标签
}
