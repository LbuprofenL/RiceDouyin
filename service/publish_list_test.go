package service

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPublishList(t *testing.T) {
	var uid int64 = DemoUser.Id
	for _, item := range DemoVideos {
		err := PublishVideo(uid, item.Title, item.VideoURL, item.CoverURL)
		assert.Equal(t, nil, err)
		time.Sleep(1 * time.Second)
	}
	ans, err := PublishList(strconv.Itoa(int(uid)), uid)
	assert.Equal(t, nil, err)
	for index, item := range *ans {
		assert.Equal(t, index, item.Id)
		assert.Equal(t, strconv.Itoa(index), item.Title)
		assert.Equal(t, true, item.IsFavorite)
	}
}
