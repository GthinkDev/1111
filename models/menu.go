package models

import (
	"gin-blog/models/ctype"
	"gorm.io/gorm"
)

// MenuInfo 菜单信息
type MenuInfo struct {
	gorm.Model
	MenuTitle    string      `json:"menu_title" gorm:"size:32"`                                                         // 菜单名称
	MenuTitleEn  string      `json:"menu_title_en" gorm:"size:32"`                                                      // 菜单英文名称
	Slogan       string      `json:"slogan" gorm:"size:32"`                                                             // 菜单 Slogan
	Abstract     ctype.Array `json:"abstract" gorm:"size:32"`                                                           // 菜单简介
	AbstractTime int         `json:"abstract_time"`                                                                     // 菜单简介的切换时间
	Images       []ImageInfo `json:"images" gorm:"many2many:menu_images; joinForeignKey:MenuID;JoinReferences:ImageID"` // 菜单图片
	ImageTime    int         `json:"image_time"`                                                                        // 菜单时间
	Sort         int         `json:"sort" gorm:"size:10"`                                                               // 菜单的排序
}
