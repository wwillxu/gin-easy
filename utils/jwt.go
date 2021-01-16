package utils

import (
	"gin-easy/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	// jwt过期时间
	jwtTimeout = time.Hour * 24 * 5
)

func GenerateToken(data interface{}) (string, error) {
	// 生成token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// 载入payload
	claims["id"] = data
	// 载入过期时间
	claims["exp"] = time.Now().Add(jwtTimeout).Unix()
	return token.SignedString([]byte(config.Conf.JWT.Secret))
}
