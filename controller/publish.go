package controller

import (
	"RiceDouyin/service"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	prefix_ip       = "http://10.3.116.18:8080/douyin/"
	ffmpegPath      = ".\\util\\ffmpeg.exe"
	VideoPathPrefix = ".\\archive\\video\\"
	CoverPathPrefix = ".\\archive\\cover\\"
	VideoSuffix     = ".mp4"
	CoverSuffix     = ".jpg"
)

type VideoListResponse struct {
	Response
	VideoList []service.VideoItem `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	// Check token
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	//Get video data
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	user := usersLoginInfo[token]
	title := c.PostForm("title")

	finalName := fmt.Sprintf("%d_%d", user.Id, time.Now().Unix()) //生成新的文件名

	videoFileName := finalName + VideoSuffix
	coverFileName := finalName + CoverSuffix
	SaveVideoPath := filepath.Join("./archive/video/", videoFileName) //组合视频保存路径和文件名
	outputImagePath := "archive/cover/" + coverFileName
	if err := c.SaveUploadedFile(data, SaveVideoPath); err != nil {
		//Save video file
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//保存视频封面文件
	err = saveCover(SaveVideoPath, outputImagePath)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	Vurl := prefix_ip + "video/" + videoFileName
	Curl := prefix_ip + "cover/" + coverFileName
	service.PublishVideo(user.Id, title, Vurl, Curl)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}
func saveCover(videoPath string, coverPath string) error {
	cmd := exec.Command(ffmpegPath, "-i", videoPath, "-y", "-vframes", "1", coverPath)
	// 获取输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// 开始执行命令
	if err := cmd.Start(); err != nil {
		return err
	}
	errChan := make(chan error)
	// 读取输出
	go func() {
		if _, err := io.Copy(os.Stdout, stdout); err != nil {
			return
		}
		errChan <- err
	}()
	<-errChan
	if errChan != nil {
		return err
	}
	// 等待命令执行完成
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

// PublishList
func PublishList(c *gin.Context) {
	token := c.PostForm("token")

	// Check token
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	authorIdStr := c.Query("user_id")
	user := usersLoginInfo[token]
	authorId, err := strconv.ParseInt(authorIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "将authorIdStr解析为int时出错"})
		return
	}

	DemoVideos, err := service.PublishList(authorId, user.Id)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Fail to get publishlist"})
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: *DemoVideos,
	})
}
