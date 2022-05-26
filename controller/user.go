package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

//// usersLoginInfo use map to store user info, and key is username+password for demo
//// user data will be cleared every time the server starts
//// test data: username=zhanglei, password=douyin
//var usersLoginInfo = map[string]User{
//	"zhangleidouyin": {
//		//Id:            1,
//		Name:          "zhanglei",
//		FollowCount:   10,
//		FollowerCount: 5,
//		IsFollow:      true,
//	},
//}
//

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// BCryptPassword 对密码加密
func BCryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	passwordHash, _ := BCryptPassword(password) // 对密码加密，不保存明文
	// 根据用户名的唯一性，查找是否存在该用户，如果不存在则将用户信息存入数据库中
	var user User
	if rowsAffected := DB.Where("name = ?", username).First(&user).RowsAffected; rowsAffected != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exists"},
		})
	} else {
		newUser := User{
			Name:     username,
			Password: passwordHash,
		}
		DB.Create(&newUser) // 存入数据库
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   newUser.Id,
			Token:    username, // 由于各个接口都依靠 Token 确定用户身份，因此返回用户名
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// 查找数据库中对应的用户名，并检查密码
	var user User
	if rowsAffected := DB.Where("name = ?", username).First(&user).RowsAffected; rowsAffected != 0 {
		passwordHashByte := []byte(user.Password)
		passwordByte := []byte(password)
		// 检查密码是否正确，使用 BCrypt 内置的比较函数
		if err := bcrypt.CompareHashAndPassword(passwordHashByte, passwordByte); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// 密码错误
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Incorrect password"},
			})
		} else { // 密码正确
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    user.Name, // 由于各个接口都依靠 Token 确定用户身份，因此返回用户名
			})
		}
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	// demo 这里判断了 user 是否存在，但个人认为不用，因此省去
	user := User{Id: userId}
	DB.First(&user)
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})
}
