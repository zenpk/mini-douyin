package controller

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`
	Author        User   `json:"author" gorm:"foreignKey:UserId"`
	UserId        int64  `gorm:"not null"` // 视频对应的用户 Id
	PlayUrl       string `json:"play_url" json:"play_url,omitempty" gorm:"not null"`
	CoverUrl      string `json:"cover_url,omitempty" gorm:"not null"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey"`
	User       User   `json:"user" gorm:"foreignKey:UserId"`
	UserId     int64  `gorm:"not null"` //评论对应的用户id
	Video      Video  `gorm:"foreignKey:VideoId"`
	VideoId    int64  `gorm:"not null"` //评论对应的视频ID
	Content    string `json:"content,omitempty" gorm:"not null"`
	CreateDate string `json:"create_date,omitempty" gorm:"not null"`
}

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`
	Name          string `json:"name,omitempty" gorm:"unique; not null"`
	Password      string `gorm:"not null"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

// Favorite 记录用户点赞的视频
type Favorite struct {
	Id      int64 `gorm:"primaryKey"`
	User    User  `gorm:"foreignKey:UserId"`
	UserId  int64 `gorm:"not null"`
	Video   Video `gorm:"foreignKey:VideoId"`
	VideoId int64 `gorm:"not null"`
}
