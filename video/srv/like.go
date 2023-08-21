package srv

import (
	"context"
	"strings"
	"time"
	"tinytiktok/utils/msg"
	"tinytiktok/utils/redis"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/like"
)

func FlushLikeDataToMysql() {
	client := redis.Client
	// 获取所有键
	keys, err := client.Keys("*").Result()
	if err != nil {
		return
	}
	// 遍历所有键并获取对应的哈希值
	for _, key := range keys {
		// 该协程只处理like类型的数据
		// TODO 后期请将类似数据整合在一起集中处理
		if !strings.Contains(key, "like") {
			continue
		}
		l := &models.Like{}
		_ = redis.GetHash(key, l)
		// 如果数据变动
		if l.IsEdit {
			_ = models.LikeVideo(VideoDb, l.UserID, l.VideoID, l.State)
			// 在刷新数据之后,将edit赋值为false, 避免重复写入数据库
			l.IsEdit = false
			_ = redis.PutHash(key, l)
		}
	}
}

func init() {
	go func() {
		for {
			time.Sleep(time.Duration(redis.RefreshTime) * time.Second)
			FlushLikeDataToMysql()
		}
	}()
}

func (h *Handle) Like(ctx context.Context, req *like.LikeRequest) (rsp *like.LikeResponse, err error) {
	rsp = &like.LikeResponse{}
	// 1 : 点赞  2 : 取消点赞
	isFavorite := req.ActionType == 1
	// 将数据同步至Redis
	err = redis.PutHash(redis.Key("like", req.UserId, req.VideoId), &models.LikeCache{
		UserID:  req.UserId,
		VideoID: req.VideoId,
		State:   isFavorite,
		Table:   "likes",
		// edit赋值为true, 表示该值已被更新, 请刷新至数据库
		IsEdit: true,
	})
	if err != nil {
		rsp.StatusCode = msg.Fail
		rsp.StatusMsg = msg.RepeatError
		return rsp, err
	}
	// 返回结果
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	return rsp, nil
}
