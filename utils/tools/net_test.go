package tools

import (
	"fmt"
	"testing"
)

func TestGetPort(t *testing.T) {
	fmt.Println(GetFreePort())
	fmt.Println(GetLocalIP())
}
