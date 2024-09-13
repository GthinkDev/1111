package models

import "gorm.io/gorm"

type SysInfo struct {
	gorm.Model
	Version   string `json:"version" gorm:"size:32"`    // 版本号
	SiteTitle string `json:"site_title" gorm:"size:32"` // 网站标题
	SiteLogo  string `json:"site_logo" gorm:"size:32"`  // 网站logo
	SiteIcp   string `json:"site_icp" gorm:"size:32"`   // 网站备案号
}
