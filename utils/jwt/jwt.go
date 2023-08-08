package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
	"tinytiktok/utils/config"
)

var (
	SigningKey       []byte
	ExpiresTime      int
	TokenValid       = "token is valid"
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type UserClaims struct {
	ID   int64
	Name string
	jwt.StandardClaims
}

func init() {
	path := os.Getenv("APP")
	jwtConfig := config.NewConfig(fmt.Sprintf("%s/config", path), "jwt.yaml", "yaml")
	SigningKey = []byte(jwtConfig.ReadString("Key"))
	ExpiresTime = jwtConfig.ReadInt("ExpiresTime")
}

// CreateToken 创建一个token
func CreateToken(claims *UserClaims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		// 签名的生效时间
		NotBefore: time.Now().Unix(),
		// 过期时间
		ExpiresAt: time.Now().Unix() + int64(ExpiresTime),
		Issuer:    "tinytiktok",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

// ParseToken 解析 token
func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

// RefreshToken 更新token
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return CreateToken(claims)
	}
	return "", TokenInvalid
}
