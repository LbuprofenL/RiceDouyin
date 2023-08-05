package service

import (
	"RiceDouyin/dao"
)

var DemoVideos = []VideoItem{
	{
		Id:            0,
		Author:        DemoUser,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "0",
		IsFavorite:    false,
	},
	{
		Id:            1,
		Author:        DemoUser,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "1",
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUser,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "2",
		IsFavorite:    false,
	},
		{
		Id:            3,
		Author:        DemoUser,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "3",
		IsFavorite:    false,
	},
}

var DemoUser = dao.User{
	Id:       1,
	Name:     "TestUser",
	IsFollow: false,
	Avatar:   "vvv",
}
