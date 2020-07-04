package views

const (
	Success = 20000

	// 客户端通用错误
	ErrorCliParam = 40000
	// 客户端登陆信息错误
	ErrCliLogin = 40001
	// 客户端请求的用户已存在
	ErrCliUserExist = 40003

	// 服务端通用错误
	ErrorServer = 50000
)

var errMap = map[int]string{
	//ErrorCliParam:     "[Param Error] only use in controller/common.go",
	ErrCliLogin:     "[Client Error] wrong user name or password",
	ErrCliUserExist: "[Client Error] user exist",
	ErrorServer:     "[Server Error] contact admin",
}

func GetErrMsg(code int) string {
	msg, err := errMap[code]
	if !err {
		return "[Undefined Error] Contact developer"
	}
	return msg
}
