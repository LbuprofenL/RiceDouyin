package dao

import (
	"errors"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	Id            int64     `gorm:"id"`             //视频id
	UserId        int64     `gorm:"user_id"`        //作者id
	Title         string    `gorm:"title"`          //视频标题
	FavoriteCount int32     `gorm:"favorite_count"` //视频获赞数
	CommentCount  int32     `gorm:"comment_count"`  //视频评论数
	VideoURL      string    `gorm:"video_url"`      //视频文件保存路径
	CoverURL      string    `gorm:"cover_url"`      //视频封面文件保存路径
	CreateTime    time.Time `gorm:"create_time"`    //视频上传时间
	UpdateTime    time.Time `gorm:"update_time"`    //记录更新时间
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

// 返回VideoDao对象,实现并发安全的惰性初始化
func NewVideoInstance() *VideoDao {
	videoOnce.Do(func() {
		videoDao = &VideoDao{}
	})
	return videoDao
}

// 添加记录
func (*VideoDao) PublishNewVideo(v *Video) error {
	if err := db.Model(&Video{}).Omit("id", "create_time", "update_time").Create(v).Error; err != nil {
		return errors.New("数据库添加视频记录时出错")
	}
	return nil
}

// 修改记录(获赞数，评论数)
func (*VideoDao) LikeVideo(vid int64) error {
	err := db.Model(&Video{}).Where("id = ?", vid).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	if err != nil {
		return errors.New("数据库添加点赞记录时出错")
	}
	return nil
}

func (*VideoDao) LikeVideoWithTransaction(tx *gorm.DB, vid int64) error {
	err := tx.Model(&Video{}).Where("id = ?", vid).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	if err != nil {
		return errors.New("数据库添加点赞记录时出错")
	}
	return nil
}

func (*VideoDao) CommentVideo(vid int64) error {
	err := db.Model(&Video{}).Where("id = ?", vid).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		return errors.New("数据库添加评论记录时出错")
	}
	return nil
}
func (*VideoDao) CommentVideoWithTransaction(tx *gorm.DB, vid int64) error {
	err := tx.Model(&Video{}).Where("id = ?", vid).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		return errors.New("数据库添加评论记录时出错")
	}
	return nil
}

// 查询记录(按时间，按视频id，按作者id)

// 返回此时间之前发布的视频数量
func (*VideoDao) GetVideoSumByTime(latest_time time.Time) (int, error) {
	var count int64
	result := db.Model(&Video{}).Where("create_time < ?", latest_time).Count(&count)
	if result.Error != nil {
		return 0, errors.New("返回此时间之前发布的视频数量出错")
	}
	return int(count), nil
}

// 根据此时间之前发布的第一条视频
func (*VideoDao) GetVideoByTime(latest_time time.Time) (*Video, error) {
	var v Video
	err := db.Model(&Video{}).Order("create_time DESC").Where("create_time < ?", latest_time).First(&v).Error
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// 根据视频id返回视频记录
func (*VideoDao) GetVideoByVideoId(vid int64) (*Video, error) {
	var v = &Video{}
	if err := db.Model(&Video{}).Where("id = ?", vid).Find(&v).Error; err != nil {
		return nil, errors.New("按照视频id查询视频时出错")
	}
	return v, nil
}

// 根据用户id返回视频数量
func (*VideoDao) GetVideoCountByUserId(uid int64) (int, error) {
	var count int64
	result := db.Model(&Video{}).Where("user_id = ?", uid).Count(&count)
	if result.Error != nil {
		return 0, errors.New("根据用户id返回视频数量时出错")
	}

	return int(count), nil
}

// 根据用户id返回指定数量的视频
func (*VideoDao) GetVideoByUserId(uid int64, count int) (*[]Video, error) {
	videos := make([]Video, count)
	result := db.Model(&Video{}).Order("create_time desc").Where("user_id = ?", uid).Limit(count).Find(&videos)
	if result.Error != nil {
		return nil, errors.New("按照用户id查询视频时出错")
	}
	return &videos, nil
}

// 是否存在指定视频id的视频
func (*VideoDao) HasVideoById(vid int64) (bool, error) {
	r := db.Model(&Video{}).Where("id = ?", vid).Limit(1).Find(&Video{})
	if r.Error != nil {
		return false, errors.New("按照视频id查询视频是否存在时出错")
	}
	exists := r.RowsAffected > 0
	return exists, nil
}

// 是否存在指定的用户id
func (*VideoDao) IsUserIdValid(uid int64) (bool, error) {
	r := db.Model(&Video{}).Where("user_id = ?", uid).Limit(1).Find(&Video{})
	if r.Error != nil {
		return false, errors.New("根据用户id返回视频是否存在时出错")
	}
	exists := r.RowsAffected > 0
	return exists, nil
}
