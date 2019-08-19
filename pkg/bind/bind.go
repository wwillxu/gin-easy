package bind

import (
	"encoding/json"
	"gineasy/pkg/e"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

func BindCheck(c *gin.Context,bindValue interface{}) int{
	err := c.ShouldBind(bindValue)
	if err!=nil{
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _ = range ve {
				return e.ErrorInputMissing
			}
		}
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return e.ErrorInputType
		}
		return e.Error
	}
	return e.Success
}