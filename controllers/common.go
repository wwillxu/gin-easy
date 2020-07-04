package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"mime/multipart"
)

var ParamError = errors.New("param error")

// 参数绑定
func reqValidator(c *gin.Context, value interface{}) error {
	c.ShouldBind(value)
	validate := validator.New()
	err := validate.Struct(value)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			unPassValidator(c, err)
			return ParamError
		}
	}
	return nil
}

// 获取当前用户
func getUserID(c *gin.Context) (primitive.ObjectID, error) {
	// 获取当前ID
	value, exists := c.Get("id")
	if !exists {
		unPassValidator(c, "user id get failed")
		return primitive.NilObjectID, ParamError
	}
	res := fmt.Sprintf("%v", value)
	// 检查ID合法性
	id, err := primitive.ObjectIDFromHex(res)
	if err != nil {
		unPassValidator(c, "user id is invalid")
		return primitive.NilObjectID, ParamError
	}
	return id, nil
}

func getParamID(c *gin.Context) (primitive.ObjectID, error) {
	// 获取当前ID
	value := c.Param("id")
	if value == "" {
		unPassValidator(c, "input nil id")
		return primitive.NilObjectID, ParamError
	}
	// 检查ID合法性
	id, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		unPassValidator(c, "user id is invalid")
		return primitive.NilObjectID, ParamError
	}
	return id, nil
}

func getQueryKey(c *gin.Context, key string) (string, error) {
	value := c.Query(key)
	if value == "" {
		unPassValidator(c, "required query key: "+key)
		return "", ParamError
	}
	return value, nil
}

func getFormFile(c *gin.Context, key string) (*multipart.FileHeader, error) {
	value, err := c.FormFile(key)
	if err != nil {
		unPassValidator(c, err)
		return nil, ParamError
	}
	return value, nil
}

func unPassValidator(c *gin.Context, err interface{}) {
	c.JSON(200, gin.H{
		"code":  40000,
		"error": fmt.Sprintf("[Param Error] %s", err),
		"data":  nil,
	})
}
