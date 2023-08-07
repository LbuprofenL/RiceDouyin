package service

import (
	"RiceDouyin/dao"
)

var DemoVideos = []VideoItem{
	{
		Id:            0,
		Author:        DemoUserA,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "aaa",
		IsFavorite:    false,
	},
	{
		Id:            1,
		Author:        DemoUserB,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "bbb",
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUserC,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "ccc",
		IsFavorite:    false,
	},
	{
		Id:            3,
		Author:        DemoUserD,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "444",
		IsFavorite:    false,
	},
	{
		Author:        DemoUserD,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "0",
		IsFavorite:    false,
	},
}

var DemoUser = dao.User{
	Id:       1,
	Name:     "TestUser",
	IsFollow: false,
	Avatar:   "vvv",
}

var DemoUserA = dao.User{
	Id:       1,
	Name:     "TestUserA",
	IsFollow: true,
}

var DemoUserB = dao.User{
	Id:       2,
	Name:     "TestUserB",
	IsFollow: false,
}
var DemoUserC = dao.User{
	Id:       3,
	Name:     "TestUserC",
	IsFollow: false,
}
var DemoUserD = dao.User{
	Id:       4,
	Name:     "TestUserD",
	IsFollow: true,
}

var DemoUserGrp = [4]dao.User{
	DemoUserA, DemoUserB, DemoUserC, DemoUserD,
}
