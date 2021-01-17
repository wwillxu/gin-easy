package context

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (string, error) {
	value, exists := c.Get("id")
	if !exists {
		return "", errors.New("no login info")
	}
	return fmt.Sprintf("%v", value), nil
}
