package dfs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 测试从gin读取文件
func TestGinGetFile(t *testing.T) {
	router := gin.Default()
	router.POST("/file", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		url, _ := UploadFile(file, 1)
		ctx.JSON(http.StatusOK, gin.H{
			"URL": url,
		})
	})
	router.Run(":5052")
}
