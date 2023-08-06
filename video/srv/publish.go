package srv

import (
	"context"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/publish"
)

func (h *Handle) Publish(ctx context.Context, req *publish.PublishRequest) (rsp *publish.PublishResponse, err error) {
	rsp = &publish.PublishResponse{}
	video := &models.Video{
		AuthorID: req.AuthorId,
		PlayURL:  req.PlayUrl,
		CoverURL: req.CoverUrl,
		Title:    req.Title,
	}
	err = models.InsertVideo(VideoDb, video)
	if err != nil {
		rsp.StatusCode = 1
		rsp.StatusMsg = "not ok"
		return rsp, err
	}
	rsp.StatusCode = 0
	rsp.StatusMsg = "ok"
	return rsp, nil
}
