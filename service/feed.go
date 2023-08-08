package service

import (
	"RiceDouyin/dao"
	"errors"
	"time"
)

const (
	_maxVideoNum = 30
	_dateFormat  = "2006-01-02 15:04:05"
)

type feedFlow struct {
	userId     int64 //使用者id
	latestTime time.Time

	videoList  []*dao.Video
	authorList []*dao.User
	videoNum   int

	packedVideoList []VideoItem
}
type PackedFeed struct {
	PackedVideoList *[]VideoItem
	LatestTime      time.Time
}

func Feed(userId int64, latestTime time.Time) (*PackedFeed, error) {
	return getFeed(userId, latestTime).Do()
}

func getFeed(uid int64, lastTime time.Time) *feedFlow {

	return &feedFlow{
		userId:     uid,
		latestTime: lastTime,
	} //堆中分配，不会引起生命周期的问题
}

func (f *feedFlow) Do() (*PackedFeed, error) {
	var err error

	//参数检查
	err = f.checkParam()
	if err != nil {
		return nil, err
	}
	//参数准备
	err = f.prepareInfo()
	if err != nil {
		return nil, err
	}
	//调用Dao层
	err = f.packInfo()
	if err != nil {
		return nil, err
	}
	return &PackedFeed{
		PackedVideoList: &f.packedVideoList,
		LatestTime:      f.latestTime,
	}, nil
}

func (f *feedFlow) checkParam() error {
	//检查latestTime是否小于现在
	if f.latestTime.After(time.Now()) {
		return errors.New("时间早于当前")
	}
	return nil
}

func (f *feedFlow) prepareInfo() error {
	// 获取此时间之前发布的视频数量
	videoNum, err := dao.NewVideoInstance().GetVideoSumByTime(f.latestTime)
	if err != nil {
		return err
	}
	// 数量判断，大于30条只返回30条，不足30条返回所有
	if videoNum > _maxVideoNum {
		videoNum = _maxVideoNum
	}

	f.videoNum = videoNum
	f.authorList = make([]*dao.User, videoNum)
	f.videoList = make([]*dao.Video, videoNum)

	for i := 0; i < videoNum; i++ {
		//获取视频信息
		f.videoList[i], err = dao.NewVideoInstance().GetVideoByTime(f.latestTime)
		if err != nil {
			return err
		}
		//根据作者id获取作者信息
		f.authorList[i], err = GetUserById(f.videoList[i].UserId, f.userId)
		if err != nil {
			return err
		}
		//更新时间
		f.latestTime = f.videoList[i].CreateTime
	}
	return nil
}

func (f *feedFlow) packInfo() error {
	f.packedVideoList = make([]VideoItem, f.videoNum)
	for idx, item := range f.videoList {
		f.packedVideoList[idx].Id = item.Id
		f.packedVideoList[idx].Author = *f.authorList[idx]
		f.packedVideoList[idx].VideoURL = item.VideoURL
		f.packedVideoList[idx].CoverURL = item.CoverURL
		f.packedVideoList[idx].Title = item.Title
		f.packedVideoList[idx].FavoriteCount = item.FavoriteCount
		f.packedVideoList[idx].CommentCount = item.CommentCount
		//TODO并发优化
		f.packedVideoList[idx].IsFavorite, _ = dao.NewFavoriteInstance().IsFavorite(f.userId, item.Id)
	}
	return nil
}
