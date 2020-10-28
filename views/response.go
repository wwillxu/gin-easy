package views

type response struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func BasicResponse(code int, err string, data interface{}) interface{} {
	return response{
		Code:  code,
		Error: err,
		Data:  data,
	}
}

func DataResponse(data interface{}) interface{} {
	return BasicResponse(Success, "", data)
}

func ErrResponse(code int) interface{} {
	return BasicResponse(code, getErrMsg(code), nil)
}
