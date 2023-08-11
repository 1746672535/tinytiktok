package srv

import (
	"context"
	"fmt"
	"testing"
	"tinytiktok/user/proto/login"
)

func TestLogin(t *testing.T) {
	// 创建一个空的上下文对象
	ctx := context.Background()
	// 构造一个注册请求对象
	req := &login.LoginRequest{
		Username: "2@qq.com",
		Password: "222222",
	}
	// 创建一个 Handle 对象
	h := &Handle{}
	// 调用 Login 方法
	rsp, err := h.Login(ctx, req)
	if err != nil {
		t.Errorf("Register failed: %v", err)
	}
	// 检查返回值是否正确
	if rsp != nil {
		fmt.Println(rsp)
	}
}
