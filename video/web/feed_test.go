package web

import (
	"fmt"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"
)

func TestFeed(t *testing.T) {
	md := metadata.Pairs(
		"name", "jiudan",
		"name-bin", "有点心急",
	)
	rsp, _ := feedSrv(md, time.Now().Unix())
	fmt.Println(rsp)
}
