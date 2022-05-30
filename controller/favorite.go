package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type FavoriteListResponse struct {
	Response
	FavoriteList []Favorite `json:"favorite_list,omitempty"`
}

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	//判断用户是否登录
	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	}
	//赞请求的具体操作
	action := c.Query("action_type")
	actionType, err := strconv.Atoi(action)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if actionType == 1 {
		AddFavorite(c)
	} else {
		DeleteFavorite(c)
	}

}

// DeleteFavorite 取消赞
func DeleteFavorite(c *gin.Context) {
	//获取用户的userId和videoId
	userid := c.Query("user_id")
	userId, _ := strconv.ParseInt(userid, 10, 64)
	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)

	//开启数据库事务，在favorites中删除记录，在videos中更改点赞数目
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id=?", userId).Where("video_id=?", videoId).Delete(&Favorite{}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		tx.Model(&Video{}).Where("id=?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1))
		// 返回 nil 提交事务
		return nil
	})

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "Successfully unliked",
	})
}

// AddFavorite 点赞
func AddFavorite(c *gin.Context) {
	//获取用户的userId和videoId
	userid := c.Query("user_id")
	userId, _ := strconv.ParseInt(userid, 10, 64)

	videoIdStr := c.Query("video_id")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)

	//在favorites添加记录
	favorite := Favorite{
		UserId:  userId,
		VideoId: videoId,
	}

	//开启数据库事务，在favorites中添加记录，在videos中更改点赞数目
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&favorite).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		tx.Model(&Video{}).Where("id=?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1))
		// 返回 nil 提交事务
		return nil
	})

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "Thumb up success",
	})
}

// FavoriteList 获取点赞视频列表
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	//判断用户是否登录
	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	}

	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)

	var favoriteList []Favorite
	var videoList []Video
	DB.Where("user_id=?", userId).Find(&favoriteList)

	DB.Table("favorites").Select("favorites.video_id,videos.*").
		Where("favorites.user_id=?", userId).
		Joins("LEFT JOIN videos ON favorites.video_id = videos.id").
		Find(&videoList)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "success"},
		VideoList: videoList,
	})
}
