package dao

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUpdateVideo(t *testing.T) {
	var vid int64 = 111
	ok, err := NewVideoInstance().HasVideoById(vid)
	assert.Equal(t, nil, err)
	if !ok {
		err := NewVideoInstance().PublishNewVideo(
			&Video{
				Id:            vid,
				FavoriteCount: 0,
				CommentCount:  0,
				CreateTime:    time.Now()})
		assert.Equal(t, nil, err)
	}

	v, err := NewVideoInstance().GetVideoByVideoId(vid)
	assert.Equal(t, nil, err)
	oldFavorCnt := v.FavoriteCount
	oldCommCnt := v.CommentCount

	err = NewVideoInstance().LikeVideo(vid)
	assert.Equal(t, nil, err)
	err = NewVideoInstance().CommentVideo(vid)
	assert.Equal(t, nil, err)

	nv, err := NewVideoInstance().GetVideoByVideoId(vid)
	assert.Equal(t, nil, err)
	newFavorCnt := nv.FavoriteCount
	newCommCnt := nv.CommentCount

	assert.Equal(t, oldCommCnt+1, newCommCnt)
	assert.Equal(t, oldFavorCnt+1, newFavorCnt)
	fmt.Println(nv)
}
