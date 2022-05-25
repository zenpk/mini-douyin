package controller

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author" gorm:"foreignKey:UserId"`
	UserId        int64  `gorm:"not null"` // 该视频作者的唯一标志
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey"`
	User       User   `json:"user" gorm:"foreignKey:UserId"`
	UserId     int64  `gorm:"not null"`           // 该评论作者的唯一标志
	Video      Video  `gorm:"foreignKey:VideoId"` // 不知是否有必要？
	VideoId    int64  `gorm:"not null"`           // 评论对应的视频 Id
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`
	Name          string `json:"name,omitempty" gorm:"unique; not null"`
	Password      string `gorm:"not null"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"` // 意义不明？是否应当另建 follows 表和 followers 表？
}
