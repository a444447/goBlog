package model

import (
	"fmt"
	"goBlog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassport,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf("fail to connect database！")
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{}, &Profile{})

	//设置数据库允许的最大连接数
	sqlDB, err := db.DB()
	if err != nil {
		log.Print("connect db server failed.")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
