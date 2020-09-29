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

func ErrorResponse(err error) interface{} {
	return basicResponse{
		Success: false,
		Error:   err.Error(),
		Data:    nil,
	}
}
