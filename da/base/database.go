package base

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"zld-jy/config"
)

var DB *gorm.DB

func InitDB() {
	var err error
	var db *gorm.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", config.Config.DB.Username, config.Config.DB.Password, config.Config.DB.Host,
		config.Config.DB.Port, config.Config.DB.Dbname, config.Config.DB.Other)
	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	DB = db
	DB.Debug()
	log.Println(">>>>>>>数据库初始化成功>>>>>>>>>")
}
