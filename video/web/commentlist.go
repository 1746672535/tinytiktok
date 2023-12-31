package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"tinytiktok/common"
	"tinytiktok/utils/consul"
	"tinytiktok/utils/msg"
	"tinytiktok/video/proto/commentList"
	"tinytiktok/video/proto/server"
)

func CommentList(ctx *gin.Context) {
	// 鉴权
	if !ctx.GetBool("auth") {
		// 无需鉴权
	}
	// 获取参数
	videoID := ctx.DefaultQuery("video_id", "-1")
	videoIDInt, err := strconv.Atoi(videoID)
	// 访问srv层
	md := metadata.Pairs()
	conn := consul.GetClientConn(common.VideoServer)
	if conn == nil {
		panic(msg.ServerFindError)
	}
	defer conn.Close()
	client := server.NewVideoServiceClient(conn)
	rsp, err := client.CommentList(metadata.NewOutgoingContext(context.Background(), md), &commentList.CommentListRequest{
		VideoId: int64(videoIDInt),
	})
	if err != nil {
		common.ReturnErr(ctx, msg.ServerError)
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code":  rsp.StatusCode,
		"status_msg":   rsp.StatusMsg,
		"comment_list": rsp.CommentList,
	})
}
