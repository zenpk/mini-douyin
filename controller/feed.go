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

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}

//func GetVideoList() {
//	var demo = []Video{
//		{
//			Id:            1,
//			Author:        DemoUser,
//			PlayUrl:       "../public/bear.mp4",
//			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
//			FavoriteCount: 0,
//			CommentCount:  0,
//			IsFavorite:    false,
//		}}
//
//}
