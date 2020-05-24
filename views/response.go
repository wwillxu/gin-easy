package views

import (
	"errors"
	"fmt"
)

var (
	ErrorClient = errors.New("客户端错误")

	ErrCliUserExist = errors.New("用户已存在")
	ErrCliLogin= errors.New("用户名或密码错误")

	ErrorServer = errors.New("服务端错误")
)

type basicResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func Response(data interface{}) interface{} {
	return basicResponse{
		Success: true,
		Error:   "",
		Data:    data,
	}
}

func ErrorResponse(err interface{}) interface{} {
	return basicResponse{
		Success: false,
		Error:   fmt.Sprint(err),
		Data:    nil,
	}
}
