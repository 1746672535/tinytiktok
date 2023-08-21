package avatar

import (
	"fmt"
	"github.com/awnumar/fastrand"
	"strconv"
)

func randomHexColor() string {
	// 生成三个随机的RGB分量
	r := fastrand.Intn(256)
	g := fastrand.Intn(256)
	b := fastrand.Intn(256)

	// 将RGB分量转换为16进制字符串
	hexR := strconv.FormatInt(int64(r), 16)
	hexG := strconv.FormatInt(int64(g), 16)
	hexB := strconv.FormatInt(int64(b), 16)

	// 如果16进制字符串只有一位，则在前面补0
	if len(hexR) == 1 {
		hexR = "0" + hexR
	}
	if len(hexG) == 1 {
		hexG = "0" + hexG
	}
	if len(hexB) == 1 {
		hexB = "0" + hexB
	}

	// 拼接RGB分量成为最终的16进制颜色
	color := hexR + hexG + hexB

	return color
}

func Gen(name string) string {
	return fmt.Sprintf("https://ui-avatars.com/api/?name=%s&background=%s&size=200", name, randomHexColor())
}
