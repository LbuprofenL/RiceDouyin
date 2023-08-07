package service

import (
	"RiceDouyin/dao"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	_times int = 8
)

func TestPublish(t *testing.T) {
	// var uid int64 = DemoUser.Id
	for _, item := range DemoVideos {
		err := PublishVideo(item.Author.Id, item.Title, item.VideoURL, item.CoverURL)
		assert.Equal(t, nil, err)
		time.Sleep(1 * time.Second)
	}
	_times += 1
}

func TestPublishList(t *testing.T) {
	TestPublishListA(t)
	TestPublishListB(t)
	TestPublishListC(t)
	// TestPublishListA(t)

}

func TestPublishListA(t *testing.T) {
	ans, err := PublishList(DemoUserA.Id, DemoUserA.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, _times, len(*ans))

	for _, item := range *ans {
		v, err := dao.NewVideoInstance().GetVideoByVideoId(item.Id)
		assert.Equal(t, nil, err)
		assert.Equal(t, v.Id, item.Id)
		assertStrEqual(t, "aaa", item.Title)
		// assert.Equal(t, .IsFavorite, item.IsFavorite)
	}
}

func TestPublishListB(t *testing.T) {
	ans, err := PublishList(DemoUserB.Id, DemoUserB.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, _times, len(*ans))

	for _, item := range *ans {
		v, err := dao.NewVideoInstance().GetVideoByVideoId(item.Id)
		assert.Equal(t, nil, err)
		assert.Equal(t, v.Id, item.Id)
		assertStrEqual(t, "bbb", item.Title)
		// assert.Equal(t, .IsFavorite, item.IsFavorite)
	}
}

func TestPublishListC(t *testing.T) {
	ans, err := PublishList(DemoUserC.Id, DemoUserC.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, _times, len(*ans))

	for _, item := range *ans {
		v, err := dao.NewVideoInstance().GetVideoByVideoId(item.Id)
		assert.Equal(t, nil, err)
		assert.Equal(t, v.Id, item.Id)
		assertStrEqual(t, "ccc", item.Title)
		// assert.Equal(t, .IsFavorite, item.IsFavorite)
	}
}
func TestPublishListD(t *testing.T) {
	ans, err := PublishList(DemoUserA.Id, DemoUserA.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, _times, len(*ans))

	for _, item := range *ans {
		// v, err := dao.NewVideoInstance().GetVideoByVideoId(item.Id)
		assert.Equal(t, nil, err)
		// assert.Equal(t, v.Id, item.Id)
		assertStrEqual(t, "aaa", item.Title)
		// assert.Equal(t, .IsFavorite, item.IsFavorite)
	}
}
func assertStrEqual(t *testing.T, want string, got string) {
	if want != got {
		t.Errorf("want %s bug got %s", want, got)
	}
}
