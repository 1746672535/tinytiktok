package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	Key = "Ar&lqr%BSOC%%xs^zDi7W@9Wv7N4#ZZv"
	j := NewJWT()
	claims := UserClaims{
		ID:   15,
		Name: "Alice",
		StandardClaims: jwt.StandardClaims{
			// 签名的生效时间
			NotBefore: time.Now().Unix(),
			// 过期时间
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "tinytiktok",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		// 服务器错误
		fmt.Println("code: 500")
	}
	user, err := j.ParseToken(token)
	if err != nil {
		fmt.Println("token 解析错误")
		return
	}
	fmt.Println(user.ID, user.Name)
}
