package service

import (
	"RiceDouyin/dao"
	"errors"
)

const (
	_maxLen    = 32
	_maxURLLen = 255
)

type publishVideoFlow struct {
	userId   int64
	title    string
	videoURL string
	coverURL string

	// videoId    int64
	// createTime time.Time
}

func PublishVideo(userId int64, title string, videoURL string, coverURL string) error {
	// 字符串解析
	return newPublishiVideo(userId, title, videoURL, coverURL).Do()
}
func newPublishiVideo(uid int64, tl string, vURL string, cURL string) *publishVideoFlow {
	return &publishVideoFlow{
		userId:   uid,
		title:    tl,
		videoURL: vURL,
		coverURL: cURL,
	}
}
func (p *publishVideoFlow) Do() error {
	var err error

	//参数检查
	err = p.checkParam()
	if err != nil {
		return err
	}
	//参数准备
	err = p.prepareInfo()
	if err != nil {
		return err
	}
	//调用Dao层
	err = p.saveInfo()
	if err != nil {
		return err
	}
	return nil
}

func (p *publishVideoFlow) checkParam() error {
	if len(p.title) > _maxLen {
		return errors.New("视频标题长度不应超过20个字符")
	}
	//开发期间添加url检查
	if len(p.videoURL) > _maxURLLen {
		return errors.New("视频文件地址长度不应超过255个字符")
	}
	if len(p.coverURL) > _maxURLLen {
		return errors.New("封面文件地址长度不应超过255个字符")
	}
	return nil
}

func (p *publishVideoFlow) prepareInfo() error {
	// node, err := util.CreateNode()
	// if err != nil {
	// 	return errors.New("创建视频id节点时失败")
	// }

	// p.videoId = node.Generate().Int64()
	// p.createTime = time.Now()
	return nil
}

func (p *publishVideoFlow) saveInfo() error {
	err := dao.NewVideoInstance().PublishNewVideo(
		&dao.Video{
			// Id:         p.videoId,
			// CreateTime: p.createTime
			UserId:   p.userId,
			Title:    p.title,
			VideoURL: p.videoURL,
			CoverURL: p.coverURL,
		})
	if err != nil {
		return err
	}
	//修改用户投稿数量
	err = dao.NewUserInstance().PublishNewVideo(p.userId)
	if err != nil {
		return err
	}
	return nil
}
