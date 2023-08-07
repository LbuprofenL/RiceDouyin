package service

import (
	"RiceDouyin/dao"
	"errors"
)

type VideoItem struct {
	Id            int64    `json:"id,omitempty"`
	Author        dao.User `json:"author,omitempty"`
	VideoURL      string   `json:"play_url,omitempty"`
	CoverURL      string   `json:"cover_url,omitempty"`
	Title         string   `json:"title,omitempty"`
	FavoriteCount int32    `json:"favorite_count,omitempty"`
	CommentCount  int32    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
}

type publishListFlow struct {
	userId    int64
	authorId  int64
	author    *dao.User
	videoList *[]dao.Video

	packedVideoList []VideoItem
}

// 查询发布列表，传入用户id用于查询关注关系和点赞关系
func PublishList(authorId int64, userId int64) (*[]VideoItem, error) {
	return getPublishList(authorId, userId).Do()
}

func getPublishList(authorId int64, userId int64) *publishListFlow {

	return &publishListFlow{
		authorId: authorId,
		userId:   userId,
	}
}

func (p *publishListFlow) Do() (*[]VideoItem, error) {
	var err error

	//参数检查
	err = p.checkParam()
	if err != nil {
		return nil,err
	}
	//参数准备
	err = p.prepareInfo()
	if err != nil {
		return nil,err
	}
	//调用Dao层
	err = p.packInfo()
	if err != nil {
		return nil,err
	}
	return &p.packedVideoList,nil
}

func (p *publishListFlow) checkParam() error {
	// TODO:检查用户id有效性
	// ok, err := dao.NewUserInstance().IsUserIdValid(p.authorId)
	ok, err := dao.NewVideoInstance().IsUserIdValid(p.authorId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("不存在此用户")
	}
	return nil
}

// TODO:并发
func (p *publishListFlow) prepareInfo() error {
	var err error
	//TODO:获取作者信息

	p.author, err = GetUserById(p.userId, p.authorId)
	if err != nil {
		return err
	}

	//TODO:获取用户发布视频数量
	// count, err := dao.NewUserInstance().GetVideoCountByUserId(p.userId)
	count, err := dao.NewVideoInstance().GetVideoCountByUserId(p.authorId)
	if err != nil {
		return err
	}
	//获取视频列表
	p.videoList, err = dao.NewVideoInstance().GetVideoByUserId(p.authorId, count)
	if err != nil {
		return err
	}
	return nil
}

func (p *publishListFlow) packInfo() error {
	//信息打包组装
	p.packedVideoList = make([]VideoItem, len(*p.videoList))
	for i := 0; i < len(*p.videoList); i++ {
		p.packedVideoList[i].Id = (*p.videoList)[i].Id
		p.packedVideoList[i].Author = *p.author
		p.packedVideoList[i].VideoURL= (*p.videoList)[i].VideoURL
		p.packedVideoList[i].CoverURL = (*p.videoList)[i].CoverURL
		p.packedVideoList[i].Title = (*p.videoList)[i].Title
		p.packedVideoList[i].FavoriteCount = (*p.videoList)[i].FavoriteCount
		p.packedVideoList[i].CommentCount = (*p.videoList)[i].CommentCount
		
		p.packedVideoList[i].IsFavorite,_ = dao.NewFavoriteInstance().IsFavorite(p.userId,(*p.videoList)[i].Id)

	}
	return nil
}
