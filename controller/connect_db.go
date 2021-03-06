package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 连接数据库，并创建各种需要的表
// 目前已完成用户信息表

var DB *gorm.DB

func ConnectDB() {
	// 连接数据库
	dsn := DBUserPass + "@tcp(localhost:3306)/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Database connection failed")
	}
	// 创建 User, Video, Comment, Favorite, Relation 表
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
	if err := DB.AutoMigrate(&Relation{}); err != nil {
		log.Fatalln("Relation table creation failed")
	}
}
