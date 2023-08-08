package service

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublishVideo(t *testing.T) {
	var uid int64 = 2
	err := PublishVideo(uid, "333", "http://10.3.116.18:8080/douyin/video/2_1691423209.mp4", "http://10.3.116.18:8080/douyin/cover/2_1691423209.jpg")
	assert.Equal(t, nil, err)

	ans, err := PublishList(uid, uid)
	assert.Equal(t, nil, err)
	for index, item := range *ans {
		assert.Equal(t, index, item.Id)
		assert.Equal(t, strconv.Itoa(index), item.Title)
	}
}
