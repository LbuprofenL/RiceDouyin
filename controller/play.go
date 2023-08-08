package controller

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PlayVideo(c *gin.Context) {
	// 打开视频文件
	path := c.Param("path")
	file, err := os.Open("./archive/video/" + path)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer file.Close()

	// 设置Content-Type为video/mp4，告诉客户端返回的是视频文件
	c.Header("Content-Type", "video/mp4")

	// 将视频文件内容复制到ResponseWriter，实现文件的传输
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func Cover(c *gin.Context) {
	// 打开视频文件
	path := c.Param("path")
	file, err := os.Open("./archive/cover/" + path)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer file.Close()

	// 设置Content-Type为video/jpg，告诉客户端返回的是视频文件
	c.Header("Content-Type", "image/jpeg")

	// 将视频文件内容复制到ResponseWriter，实现文件的传输
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
