package main

import (
	"fmt"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestInfo(t *testing.T) {
	md := metadata.Pairs(
		"name", "jiudan",
		"name-bin", "有点心急",
	)
	rsp, _ := infoSrv(md, 2)
	fmt.Println(rsp)
}
