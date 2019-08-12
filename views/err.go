package views

import (
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v8"
)

func ErrorResponse(err string) *Response {
	return &Response{
		Success:false,
		Error: err,
		Data:nil,
	}
}

func ErrorBind(err error) *Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return ErrorResponse("缺少必要数据("+fmt.Sprintf("%s %s", e.Field, e.Tag)+")")
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ErrorResponse("错误JSON类型("+fmt.Sprint(err)+")")
	}
	return ErrorResponse("参数错误("+fmt.Sprint(err)+")")
}