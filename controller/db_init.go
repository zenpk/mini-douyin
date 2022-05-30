package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

/**
数据库初始化
*/
func Init() error {
	var err error
	dsn := "root:lixinyue@tcp(127.0.0.1:3306)/community?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Database connection failed")
	}
	// 创建 User, Video, Comment, Favorite 表
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatalln("User table creation failed")
	}
	if err := DB.AutoMigrate(&Video{}); err != nil {
		log.Fatalln("Video table creation failed")
	}
	if err := DB.AutoMigrate(&Comment{}); err != nil {
		log.Fatalln("Comment table creation failed")
	}
	if err := DB.AutoMigrate(&Favorite{}); err != nil {
		log.Fatalln("Favorite table creation failed")
	}
	return err
}
