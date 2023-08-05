package controller

import (
	"RiceDouyin/service"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []service.VideoItem `json:"video_list,omitempty"`
	NextTime  int64               `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	latestTimeStr := c.Query("latest_time")
	var latestTime time.Time
	if len(latestTimeStr) == 0 {
		latestTime = time.Now()
	} else {
		l, _ := strconv.ParseInt(latestTimeStr, 10, 64)
		latestTime = time.Unix(l, 0)
	}
	token := c.Query("token")
	user := usersLoginInfo[token]

	DemoFeed, err := service.Feed(user.Id, latestTime)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	if DemoFeed == nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  errors.New("视频流指针为空").Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: *DemoFeed.PackedVideoList,
		NextTime:  DemoFeed.LatestTime.Unix(),
	})
}
