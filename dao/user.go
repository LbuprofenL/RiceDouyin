package dao

import "sync"

type User struct {
	Id       int64
	Name     string
	IsFollow bool
	Avatar   string
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

//TODO:获取视频数量
func (*UserDao) GetVideoCountByUserId(vid int64) (int, error) {
	return 0, nil
}

//TODO:检查用户id有效性
func (*UserDao) IsUserIdValid(uid int64) (bool, error) {
	return true, nil
}

// TODO:添加投稿数量
func (*UserDao) PublishNewVideo(uid int64) error {
	return  nil
}
