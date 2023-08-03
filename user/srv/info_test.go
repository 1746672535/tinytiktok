package srv

import (
	"context"
	"fmt"
	"testing"
	"tinytiktok/user/proto/info2"
)

func TestInfo(t *testing.T) {
	// 创建一个空的上下文对象
	ctx := context.Background()
	// 构造一个注册请求对象
	req := &info2.UserRequest{
		UserId: 2,
	}
	// 创建一个 Handle 对象
	h := &Handle{}
	// 调用 Login 方法
	rsp, err := h.Info(ctx, req)
	if err != nil {
		t.Errorf("Register failed: %v", err)
	}
	// 检查返回值是否正确
	if rsp != nil {
		fmt.Println(rsp)
	}
}
