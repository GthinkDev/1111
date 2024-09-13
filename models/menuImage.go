package models

// MenuImage 菜单图片
type MenuImage struct {
	MenuID    uint      `json:"menu_id"`
	MenuInfo  MenuInfo  `json:"menu_info" gorm:"foreignKey:MenuID"` // 菜单信息
	ImageID   uint      `json:"image_id"`                           // 图片ID
	ImageInfo ImageInfo `gorm:"foreignKey:ImageID"`                 // 图片信息
	Sort      int       `json:"sort" gorm:"size:10"`                // 排序
}
