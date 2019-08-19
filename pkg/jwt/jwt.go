package jwt

import (
	"gineasy/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SigningMethod = jwt.SigningMethodHS256
	Key           = []byte(conf.JwtKey)
	PayloadMsg    = "id"
	timeout       = 3 * time.Hour
)

// GenerateToken can use to get a jwt token.
func GenerateToken(data interface{}) (string, error) {
	// create token
	token := jwt.New(SigningMethod)
	claims := token.Claims.(jwt.MapClaims)
	// payload
	claims[PayloadMsg] = data
	// expire time
	expire := time.Now().UTC().Add(timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Unix()
	tokenString, err := token.SignedString(Key)
	return tokenString, err
}
