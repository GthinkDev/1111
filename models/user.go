package models

import (
	"gin-blog/models/ctype"
	"gorm.io/gorm"
)

// UserInfo 用户表
type UserInfo struct {
	gorm.Model
	NickName       string           `json:"nickname" gorm:"size:64"`                                                    // 昵称
	Username       string           `json:"user_name"  gorm:"size:64"`                                                  // 用户名
	Password       string           `json:"password"  gorm:"size:42"`                                                   // 密码
	Avatar         string           `json:"avatar_id" gorm:"size:256"`                                                  // 头像
	Email          string           `json:"email" gorm:"size:256"`                                                      // 邮箱
	Phone          string           `json:"phone"  gorm:"size:18"`                                                      // 手机号
	Addr           string           `json:"addr" gorm:"size:256"`                                                       // 地址
	Token          string           `json:"token" gorm:"size:256"`                                                      // token
	IP             string           `json:"ip" gorm:"size:32"`                                                          // ip
	Role           ctype.Role       `json:"role" gorm:"size:4; default:1"`                                              // 角色权限 1 管理员  2 普通用户 3 游客 4 禁言用户
	SignStatus     ctype.SignStatus `json:"sign_status" gorm:"type=smallint(6)" json:"sign_status"`                     // 注册方式
	ArticleInfo    []ArticleInfo    `json:"-" gorm:"ForeignKey:UserID"`                                                 // 一对多 发布的文章
	CollectsModels []ArticleInfo    `json:"-" gorm:"many2many:collects;joinForeignKey:UserID;joinReferences:ArticleID"` // 收藏的文章
}
