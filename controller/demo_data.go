package controller

var DemoVideos = []Video{
	{
		Id:            0,
		Author:        DemoUser,
		PlayUrl:       ServerAddr + "/static/videos/bear.mp4", // 把 demo 的视频
		CoverUrl:      ServerAddr + "/static/covers/bear.jpg", // 和封面换成了本地的
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         0,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            0,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
