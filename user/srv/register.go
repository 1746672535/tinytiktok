package main

import (
	"context"
	"fmt"
	"tinytiktok/user/proto/register"
)

// Register TODO
func (h *Handle) Register(ctx context.Context, req *register.RegisterRequest) (rsp *register.RegisterResponse, err error) {
	fmt.Println(req.Username, req.Password)
	return rsp, nil
}
