package controller

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       IP_PORT + "/static/videos/bear.mp4",
		CoverUrl:      IP_PORT + "/static/covers/bear.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
