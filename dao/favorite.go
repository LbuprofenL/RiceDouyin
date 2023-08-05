package dao

import (
	"sync"
	"time"
)

type Favorite struct {
	Id         int64
	UserId     int64
	VideoId    int64
	CreateTime time.Time
}

func (Favorite) TableName() string {
	return "favorite"
}

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteInstance() *FavoriteDao {
	favoriteOnce.Do(func() {
		favoriteDao = &FavoriteDao{}
	})
	return favoriteDao
}

//TODO:返回视频是否已经被点赞
func (f *FavoriteDao) IsFavorite(userId int64, videoId int64) (bool, error) {
	//错误信息报错在日志中
	return true, nil
}
