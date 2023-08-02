package tools

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	psd := "youdianxinji"
	encodePsd, salt := GenerateMd5(psd)
	fmt.Println(encodePsd)
	fmt.Println(VerifyMd5(encodePsd, psd, salt))
}
