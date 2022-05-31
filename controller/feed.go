package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed 推送最新的 30 个视频
func Feed(c *gin.Context) {
	token := c.Query("token") // token 是用户名
	var user User
	DB.Where("name=?", token).First(&user)
	var videoList []Video
	DB.Preload("Author").Order("id desc").Limit(30).Find(&videoList)
	for i, v := range videoList {
		// 查找是否存在一条当前用户给该视频点赞的记录
		// 不知道为什么这里用 RowsAffected 不行
		favorite := Favorite{Id: 0}
		DB.Where("video_id=?", v.Id).Where("user_id=?", user.Id).First(&favorite)
		if favorite.Id != 0 { // 查找到了点赞记录
			videoList[i].IsFavorite = true
		}
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
