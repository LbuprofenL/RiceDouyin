package service

import "RiceDouyin/dao"

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

//TODO:返回用户指针
func GetUserById(visitId int64,visitedId int64)(*dao.User,error){
	return &DemoUser,nil
}