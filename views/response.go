package views

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

func ErrorResponse(code int) interface{} {
	return basicResponse{
		Success: false,
		Error:   getErrMsg(code),
		Data:    nil,
	}
}
