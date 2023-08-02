package main

import (
	"fmt"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestLogin(t *testing.T) {
	username := "Jiudan"
	password := "a/d87$*+#Wq79@"
	md := metadata.Pairs(
		"name", "jiudan",
		"name-bin", "有点心急",
	)
	rsp, _ := loginSrv(md, username, password)
	fmt.Println(rsp)
}
