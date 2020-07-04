package views

type basicResponse struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func Response(data interface{}) interface{} {
	return basicResponse{
		Code:  20000,
		Error: "",
		Data:  data,
	}
}

func ErrorResponse(code int) interface{} {
	return basicResponse{
		Code:  code,
		Error: GetErrMsg(code),
		Data:  nil,
	}
}
