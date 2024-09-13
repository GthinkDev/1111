package flag

import (
	"gin-blog/global"
	"gin-blog/models"
)

// MakeMigration 初始化数据库,迁移数据库等
func MakeMigration() {
	var err error
	global.DB.SetupJoinTable(&models.UserInfo{}, "CollectsInfo", &models.Collects{})
	global.DB.SetupJoinTable(&models.MenuInfo{}, "ImageInfo", &models.MenuImage{})
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserInfo{},    // 用户表
		&models.Collects{},    // 收藏表
		&models.MessageInfo{}, // 消息表
		&models.MenuImage{},   // 菜单图片表
		&models.MenuInfo{},    // 菜单表
		&models.MoodInfo{},    // 心情表
		&models.TagInfo{},     // 标签表
		&models.SysInfo{},     // 系统配置表
		&models.ArticleInfo{}, // 文章表
		&models.AdvertInfo{},  // 广告表
		&models.CommentInfo{}, // 评论表
		&models.FadeBack{},    // 反馈表
		&models.ImageInfo{},   // 图片表
		&models.LogInfo{},     // Log表
		&models.LoginData{},   // 登录日志表
	)
	if err != nil {
		global.Log.Error("数据库迁移失败", err)
		return
	} else {
		global.Log.Info("数据库迁移成功")
	}

}
