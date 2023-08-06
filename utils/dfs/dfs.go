package dfs

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/levigross/grequests"
	"mime/multipart"
	"os"
	"strings"
	"tinytiktok/utils/config"
)

var FileServerURL string
var SavePath string

func init() {
	// 初始化配置文件
	path := os.Getenv("APP")
	dfsConfig := config.NewConfig(fmt.Sprintf("%s\\config", path), "dfs.yaml", "yaml")
	host := dfsConfig.ReadString("Host")
	port := dfsConfig.ReadInt("Port")
	FileServerURL = fmt.Sprintf("http://%s:%d", host, port)
	SavePath = dfsConfig.ReadString("SavePath")
}

func UploadFile(file *multipart.FileHeader, userID int64) (string, error) {
	fd, _ := file.Open()
	videoUUID := uuid.NewString()
	videoUUID = strings.Replace(videoUUID, "-", "", -1)
	data := &grequests.RequestOptions{
		// 无需解析视频类型
		Files: []grequests.FileUpload{{FileContents: fd, FileName: fmt.Sprintf("%s", videoUUID)}},
		Params: map[string]string{
			"output": "json",
			"scene":  "keygo",
			"path":   fmt.Sprintf("%s/%d", SavePath, userID),
		}}
	rsp, err := grequests.Post(fmt.Sprintf("%s/group1/upload", FileServerURL), data)
	if err != nil || !rsp.Ok {
		return "", err
	}
	return fmt.Sprintf("%s/group1/tinytiktok/video/%d/%s", FileServerURL, userID, videoUUID), err
}
