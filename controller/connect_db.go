package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 连接数据库，并创建各种需要的表
// 目前已完成用户信息表

var DB *gorm.DB

//var userIdSequence int64
//var videoIdSequence int64

func ConnectDB() {
	// 连接数据库
	dsn := DBUserPass + "@tcp(localhost:3306)/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
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
	//// 读取 User, Video 的最新 Id，如果表为空则 Id 从 1 开始
	//var lastUser User
	//var lastVideo Video
	//if err := DB.Last(&lastUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	//	userIdSequence = 0
	//} else {
	//	userIdSequence = lastUser.Id
	//}
	//if err := DB.Last(&lastVideo).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	//	videoIdSequence = 0
	//} else {
	//	videoIdSequence = lastVideo.Id
	//}
}
