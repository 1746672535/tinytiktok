package srv

import (
	"fmt"
	"testing"
)

func TestLike(t *testing.T) {
	// 取消点赞视频
	_ = LikeVideo(1, 1, false)
	// 判断视频是否被用户点赞
	fmt.Println(IsUserLikedVideo(1, 1))
	// 点赞视频
	_ = LikeVideo(1, 1, true)
	// 判断视频是否被用户点赞
	fmt.Println(IsUserLikedVideo(1, 1))
	// 获取视频点赞总数
	fmt.Println(GetVideoLikesCount(1))
}
