package middlewares

import (
	"errors"
	"gin-easy/config"
	"gin-easy/views"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type MapClaims map[string]interface{}

var (
	// ErrExpiredToken indicates JWT token has expired. Can't refresh.
	ErrExpiredToken = errors.New("token is expired")

	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = errors.New("auth header is empty")

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = errors.New("auth header is invalid")

	// ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512, RS256, RS384 or RS512
	ErrInvalidSigningAlgorithm = errors.New("invalid signing algorithm")
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtFun(c)
	}
}

func jwtFun(c *gin.Context) {
	claims, err := GetClaimsFromJWT(c)
	if err != nil {
		unauthorized(c, err.Error())
		return
	}

	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		unauthorized(c, ErrExpiredToken.Error())
		return
	}

	c.Set("id", claims["id"])

	c.Next()
}

func GetClaimsFromJWT(c *gin.Context) (MapClaims, error) {
	token, err := ParseToken(c)

	if err != nil {
		return nil, err
	}

	claims := MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	return claims, nil
}

func ParseToken(c *gin.Context) (*jwt.Token, error) {
	token, err := jwtFromHeader(c, "Authorization")
	if err != nil {
		return nil, err
	}

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}
		return []byte(config.JwtKey), nil
	})
}

func jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)
	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrInvalidAuthHeader
	}
	return parts[1], nil
}

// 错误信息返回
func unauthorized(c *gin.Context, err string) {
	c.JSON(http.StatusOK, errResponse(err))
	c.Abort()
}

func errResponse(err string) interface{} {
	return views.BasicResponse(4030, "[Auth Error] "+err, nil)
}
