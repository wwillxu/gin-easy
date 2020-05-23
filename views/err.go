package views

const (
	// 基本响应
	Success     = 200
	ErrorClient = 400
	ErrorServer = 500
)

var errMap = map[int]string{
	// 基本响应
	Success:     "",
	ErrorClient: "客户端错误",
	ErrorServer: "服务端错误",
}

func getErrMsg(code int) string {
	msg, err := errMap[code]
	if !err {
		return "未知错误"
	}
	return msg
}
