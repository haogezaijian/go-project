package config

import (
	"go-project/global"
	"gorm.io/gorm"
	"log"
	"time"
)
import "gorm.io/driver/mysql"

func initDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("数据库连接失败: " + err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal("数据库连接失败：" + err.Error())
	}

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenconns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.Db = db
}
