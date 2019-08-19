package views

import "gineasy/pkg/e"

// basic response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ErrRes(code int) *Response {
	return &Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: "Null",
	}
}

func SuccessRes(data interface{}) *Response {
	return &Response{
		Code: 200,
		Msg:  "ok",
		Data: data,
	}
}
