package srv

import (
	"context"
	"time"
	"tinytiktok/user/proto/info2"
	"tinytiktok/utils/msg"
	"tinytiktok/utils/redis"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/comment"
	"tinytiktok/video/proto/commentList"
)

//// 将评论数据flush到数据库
//func FlushCommentDataToMysql() {
//	client := redis.Client
//	// 获取所有键
//	keys, err := client.Keys("*").Result()
//	if err != nil {
//		return
//	}
//	// 遍历所有键并获取对应的哈希值
//	for _, key := range keys {
//		// 该协程只处理comment类型的数据
//		// TODO 后期请将类似数据整合在一起集中处理
//		if !strings.Contains(key, "comment") {
//			continue
//		}
//		c := &models.CommentCache{}
//		_ = redis.GetHash(key, c)
//		// 如果数据变动
//		if c.IsEdit {
//			_, _ = models.CommentVideo(VideoDb, c)
//			// 在刷新数据之后,将edit赋值为false, 避免重复写入数据库
//			_ = redis.HSet(key, "IsEdit", false)
//		}
//	}
//}

func (h *Handle) Comment(ctx context.Context, req *comment.CommentRequest) (rsp *comment.CommentResponse, err error) {
	rsp = &comment.CommentResponse{}
	// 发表评论
	if req.ActionType == 1 {

		// 获取当前时间
		currentTime := time.Now()
		// 将时间格式化为 "MM-DD" 格式
		currentDate := currentTime.Format("01-02")

		c := &models.Comment{
			UserID:    req.UserId,
			VideoID:   req.VideoId,
			Content:   req.Content,
			CreatedAt: currentDate,
		}
		var CommentID, err = models.CommentVideo(VideoDb, c)
		// 查询视频作者信息
		user, err := models.GetUserInfo(req.UserId)

		// 将数据同步至Redis
		err = redis.ZAdd(redis.Key("video", "comment", req.VideoId), []any{&models.CommentCache{
			CommentID: CommentID,
			UserID:    req.UserId,
			VideoID:   req.VideoId,
			Content:   req.Content,
			CreatedAt: currentDate,
			// edit赋值为false, 该值先进入数据库，所以无需修改
			IsEdit: false,
		}})
		//将用户信息也同步到redis
		err = redis.Set(redis.Key("user", user.Id), user)

		if err != nil {
			return rsp, err
		}
		rsp.Comment = &comment.Comment{
			Id:         CommentID,
			User:       user,
			Content:    req.Content,
			CreateDate: currentDate,
		}
	}
	// 删除评论
	if req.ActionType == 2 {
		//在mysql中删除评论
		err = models.DeleteComment(VideoDb, req.CommentId)
		//同步到redis
		err = redis.ZRem(redis.Key("video", "comment", req.VideoId), &models.CommentCache{
			CommentID: req.CommentId,
			UserID:    req.UserId,
			VideoID:   req.VideoId,
			Content:   req.Content,
			IsEdit:    false,
		})
	}

	// 返回结果
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	return rsp, nil
}

func (h *Handle) CommentList(ctx context.Context, req *commentList.CommentListRequest) (rsp *commentList.CommentListResponse, err error) {
	rsp = &commentList.CommentListResponse{}
	var Comments []*models.CommentCache
	var commentList []*comment.Comment
	err = redis.Client.ZRevRange(redis.Key("video", "comment", req.VideoId), 0, -1).ScanSlice(&Comments)
	if err != nil {
		rsp.StatusCode = msg.Fail
		rsp.StatusMsg = msg.RepeatError
		return rsp, err
	}
	for _, c := range Comments {
		var user *info2.User
		err := redis.GetHash(redis.Key("user", c.UserID), &user)
		if err != nil {
			continue
		}
		commentList = append(commentList, &comment.Comment{
			Id:         c.CommentID,
			User:       user,
			Content:    c.Content,
			CreateDate: c.CreatedAt,
		})
	}

	//// 获取所有的comments
	//comments, err := models.GetCommentListByVideoID(VideoDb, req.VideoId)
	//if err != nil {
	//	rsp.StatusCode = 1
	//	rsp.StatusMsg = msg.NotOk
	//	return rsp, err
	//}
	//
	//// 整理返回
	//var commentList []*comment.Comment
	//for _, c := range comments {
	//	// 查询视频作者信息
	//	user, err := models.GetUserInfo(c.UserID)
	//	if err != nil {
	//		continue
	//	}
	//	commentList = append(commentList, &comment.Comment{
	//		Id:      c.ID,
	//		User:    user,
	//		Content: c.Content,
	//	})
	//}

	// 返回结果
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	rsp.CommentList = commentList
	return rsp, nil
}
