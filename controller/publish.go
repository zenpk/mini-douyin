package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish 前端传入视频、token
func Publish(c *gin.Context) {
	username := c.PostForm("token") // 这里的 Token 是用户名
	// demo 里检查了用户是否存在，个人感觉没必要

	// 读取视频
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)

	var user User // 根据用户名查找用户
	DB.Where("name=?", username).First(&user)
	// 因为存储的文件名需要包含 videoId，所以先保存到数据库，利用 Id 自增特性获取 videoId
	// 相关信息存入数据库
	playUrl := ServerAddr + "/static/videos/"
	coverUrl := ServerAddr + "/static/covers/"
	video := Video{
		Author:   user,
		UserId:   user.Id,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	DB.Create(&video)
	finalName := fmt.Sprintf("%d_%d_%s", user.Id, video.Id, filename) // 保存的文件名，为防止文件名冲突，增加一项 videoId
	saveFile := filepath.Join("./public/videos/", finalName)
	// 将视频存入本地
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
	// 调用 ffmpeg 获取封面（第一帧老是黑屏，所以这里获取第 300 帧）
	// 当然更好的实践是先读取总共有多少帧，再获取中间的某一帧，这里为了简便实现就先这样了
	cmd := exec.Command(
		"ffmpeg", "-i", "./public/videos/"+finalName,
		"-vf", "select=eq(n\\, 300)", "-frames", "1",
		"./public/covers/"+finalName+".jpg",
	)
	//cmd.Stderr = os.Stderr // 输出错误信息
	if err := cmd.Run(); err != nil {
		log.Fatalln("Video cover generation failed")
	}
	// 更新数据库中的视频和封面链接
	video.PlayUrl += finalName
	video.CoverUrl += finalName + ".jpg"
	DB.Save(&video)

}

// PublishList 显示当前用户投稿的所有视频，目前的实现方式是直接从所有视频的表里根据 user_id 选取
// 效率可能较低？
func PublishList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	var videoList []Video
	DB.Preload("Author").Where("user_id=?", userId).Find(&videoList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
