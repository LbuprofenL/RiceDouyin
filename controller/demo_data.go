package controller

import (
	"RiceDouyin/dao"
	"RiceDouyin/service"
)

var DemoVideos = []service.VideoItem{
	{
		Id:            1,
		Author:        DemoUserA,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 1,
		CommentCount:  1,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUserB,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 2,
		CommentCount:  2,
		IsFavorite:    true,
	},
	{
		Id:            3,
		Author:        DemoUserC,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 3,
		CommentCount:  3,
		IsFavorite:    true,
	},
	{
		Id:            5,
		Author:        DemoUserD,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 5,
		CommentCount:  5,
		IsFavorite:    false,
	},
	{
		Id:            5,
		Author:        DemoUserB,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 5,
		CommentCount:  5,
		IsFavorite:    false,
	},
	{
		Id:            6,
		Author:        DemoUserB,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 6,
		CommentCount:  6,
		IsFavorite:    false,
	},
	{
		Id:            7,
		Author:        DemoUserC,
		VideoURL:      "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 7,
		CommentCount:  7,
		IsFavorite:    false,
	},
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
