package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction 关注取关操作
func RelationAction(c *gin.Context) {
	token := c.Query("token") // 用户名
	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	}
	actionStr := c.Query("action_type") // 1:关注，2:取关
	action, _ := strconv.Atoi(actionStr)
	if action == 1 {
		Follow(c)
	} else {
		Unfollow(c)
	}
}

func Follow(c *gin.Context) {
	userIdStr := c.Query("user_id")      // UserA
	userIdToStr := c.Query("to_user_id") // UserB
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	userIdTo, _ := strconv.ParseInt(userIdToStr, 10, 64)
	relation := Relation{
		UserAId: userId,
		UserBId: userIdTo,
	}
	// 开启数据库事务，在 relations 中添加记录，在 users 中更改关注数
	DB.Transaction(func(tx *gorm.DB) error {
		// 创建关注关系
		if err := tx.Create(&relation).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		// 增加关注数、被关注数
		tx.Model(&User{}).Where("id=?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1))
		tx.Model(&User{}).Where("id=?", userIdTo).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1))
		// 返回 nil 提交事务
		return nil
	})
	c.JSON(http.StatusOK, Response{
		StatusCode: 1,
		StatusMsg:  "Successfully followed",
	})
}

func Unfollow(c *gin.Context) {
	userIdStr := c.Query("user_id")      // UserA
	userIdToStr := c.Query("to_user_id") // UserB
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	userIdTo, _ := strconv.ParseInt(userIdToStr, 10, 64)
	// 开启数据库事务，在 relations 中添加记录，在 users 中更改关注数
	DB.Transaction(func(tx *gorm.DB) error {
		// 删除关注关系
		if err := tx.Where("user_a_id=?", userId).Where("user_b_id=?", userIdTo).Delete(&Relation{}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		// 减少关注数、被关注数
		tx.Model(&User{}).Where("id=?", userId).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1))
		tx.Model(&User{}).Where("id=?", userIdTo).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1))
		// 返回 nil 提交事务
		return nil
	})
	c.JSON(http.StatusOK, Response{
		StatusCode: 1,
		StatusMsg:  "Successfully unfollowed",
	})
}

// FollowList 展示查询用户的关注列表
func FollowList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	var followList []Relation
	// 加载 UserB 即加载当前用户关注的用户
	DB.Preload("UserB").Where("user_a_id=?", userId).Find(&followList)
	// 这里直接暴力复制了，不知道 Go 语言有无更好的方法可以提取结构体数组中的元素
	followUserList := make([]User, len(followList))
	for i, f := range followList {
		followUserList[i] = f.UserB
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: followUserList,
	})
}

// FollowerList 展示查询用户的粉丝列表
func FollowerList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	var followList []Relation
	// 加载 UserA 即加载当前用户的粉丝
	DB.Preload("UserA").Where("user_b_id=?", userId).Find(&followList)
	// 这里直接暴力复制了，不知道 Go 语言有无更好的方法可以提取结构体数组中的元素
	followUserList := make([]User, len(followList))
	for i, f := range followList {
		followUserList[i] = f.UserA
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: followUserList,
	})
}
