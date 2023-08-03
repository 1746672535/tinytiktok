package srv

import (
	"fmt"
	"testing"
	"time"
)

func TestGetVideoList(t *testing.T) {
	for _, v := range GetVideoList(time.Now().Unix()) {
		fmt.Println(v)
		fmt.Println(getUserInfo(v.AuthorId))
	}
}
