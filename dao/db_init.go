package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:00000000@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name                                                                        // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
