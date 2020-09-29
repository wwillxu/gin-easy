package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"mime/multipart"
)

// 参数绑定
func reqValidator(c *gin.Context, value interface{}) error {
	c.ShouldBind(value)
	validate := validator.New()
	err := validate.Struct(value)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err.(error)
		}
	}
	return nil
}

// 获取当前用户
func getUserID(c *gin.Context) (primitive.ObjectID, error) {
	// 获取当前ID
	value, exists := c.Get("id")
	if !exists {
		return primitive.NilObjectID, errors.New("[Param Error] user id is nil")
	}
	res := fmt.Sprintf("%v", value)
	// 检查ID合法性
	id, err := primitive.ObjectIDFromHex(res)
	if err != nil {
		return primitive.NilObjectID, errors.New("[Param Error] user id is invalid")
	}
	return id, nil
}

func getParamID(c *gin.Context) (primitive.ObjectID, error) {
	// 获取当前ID
	value := c.Param("id")
	if value == "" {
		return primitive.NilObjectID, errors.New("[Param Error] id is nil")
	}
	// 检查ID合法性
	id, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return primitive.NilObjectID, errors.New("[Param Error] id is invalid")
	}
	return id, nil
}

func getQueryKey(c *gin.Context, key string) (string, error) {
	value := c.Query(key)
	if value == "" {
		return "", errors.New("[Param Error] required query key: " + key)
	}
	return value, nil
}

func getFormFile(c *gin.Context, key string) (*multipart.FileHeader, error) {
	value, err := c.FormFile(key)
	if err != nil {
		return nil, errors.New("[Param Error] required form key: " + key)
	}
	return value, nil
}
