package service

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublishVideo(t *testing.T) {
	var uid int64 = DemoUser.Id
	for _, item := range DemoVideos {
		err := PublishVideo(uid, item.Title, item.VideoURL, item.CoverURL)
		assert.Equal(t, nil, err)
	}
	ans, err := PublishList(strconv.Itoa(int(uid)), uid)
	assert.Equal(t, nil, err)
	for index, item := range *ans {
		assert.Equal(t, index, item.Id)
		assert.Equal(t, strconv.Itoa(index), item.Title)
	}
}
