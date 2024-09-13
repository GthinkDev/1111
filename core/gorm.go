package core

import (
	"fmt"
	"gin-blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" { // 判断是否配置了数据库
		global.Log.Warnln("还未配置数据库信息，取消 gorm 连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Warnln(fmt.Sprintf("[%s] 连接数据库失败!", dsn))
	} else {
		global.Log.Infoln("数据库连接成功~~")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)              // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 设置了连接可复用的最大时间
	return db
}
