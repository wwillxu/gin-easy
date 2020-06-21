package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(data interface{}, key []byte, timeout time.Duration) (string, error) {
	// 生成token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// 载入payload
	claims["id"] = data
	// 载入过期时间
	claims["exp"] = time.Now().Add(timeout).Unix()
	// 签名
	tokenString, err := token.SignedString(key)
	return tokenString, err
}
