package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

// 请求参数绑定校验
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
func currentUser(c *gin.Context) (primitive.ObjectID, error) {
	// 获取当前ID
	value, exists := c.Get("id")
	if !exists {
		return primitive.ObjectID{}, errors.New(" User id get failed")
	}
	res := fmt.Sprintf("%v", value)
	// 检查ID合法性
	id, err := primitive.ObjectIDFromHex(res)
	if err != nil {
		log.Println(err)
		return primitive.ObjectID{}, errors.New(" User id invalid")
	}
	return id, nil
}
