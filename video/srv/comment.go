package srv

import (
	"context"
	"tinytiktok/utils/msg"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/comment"
	"tinytiktok/video/proto/commentlist"
)

func (h *Handle) Comment(ctx context.Context, req *comment.CommentRequest) (rsp *comment.CommentResponse, err error) {
	rsp = &comment.CommentResponse{}
	//发表评论
	if req.ActionType == 1 {
		c := &models.Comment{
			VideoID: req.VideoId,
			Content: req.Content,
		}
		err = models.CommentVideo(VideoDb, c)
		// 为视频的评论数量+1
		err = models.CalcCommentCountByVideoID(VideoDb, req.VideoId, true)
		if err != nil {
			rsp.StatusCode = 1
			rsp.StatusMsg = "not ok"
			return rsp, err
		}
	}
	//删除评论
	if req.ActionType == 2 {
		err = models.DeleteComment(VideoDb, req.CommentId)
		// 为视频的评论数量-1
		err = models.CalcCommentCountByVideoID(VideoDb, req.VideoId, false)
		if err != nil {
			rsp.StatusCode = 1
			rsp.StatusMsg = "not ok"
			return rsp, err
		}
	}

	// 返回结果
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	return rsp, nil
}

func (h *Handle) CommentList(ctx context.Context, req *commentlist.CommentListRequest) (rsp *commentlist.CommentListResponse, err error) {
	rsp = &commentlist.CommentListResponse{}

	// 获取所有的comments
	comments, err := models.GetCommentListByVideoID(VideoDb, req.VideoId)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = "not ok"
		return rsp, err
	}

	// 整理返回
	var commentList []*comment.Comment
	for _, c := range comments {
		// 查询视频作者信息
		user, err := models.GetUserInfo(c.UserID)
		if err != nil {
			continue
		}
		commentList = append(commentList, &comment.Comment{
			Id:      c.ID,
			User:    user,
			Content: c.Content,
		})
	}

	// 返回结果
	rsp.StatusCode = msg.Success
	rsp.StatusMsg = msg.Ok
	rsp.CommentList = commentList
	return rsp, nil
}
