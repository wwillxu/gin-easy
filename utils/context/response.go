package context

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response{
		Success: true,
		Error:   "",
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, response{
		Success: false,
		Error:   err.Error(),
		Data:    nil,
	})
}
