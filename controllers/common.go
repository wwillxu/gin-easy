package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func reqValidator(c *gin.Context, value interface{}) validator.FieldError {
	c.ShouldBind(value)
	validate := validator.New()
	err := validate.Struct(value)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}

// 获取当前用户
func currentUser(c *gin.Context) (string, error) {
	value, exists := c.Get("id")
	if !exists {
		return "", errors.New(" User id get failed")
	}
	res := fmt.Sprintf("%v", value)
	return res, nil
}
