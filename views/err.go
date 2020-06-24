package views

const (
	Success = 20000

	// 客户端通用错误
	ErrorClient = 40000
	// 客户端参数校验错误
	ErrCliParam = 40001
	// 客户端请求的用户已存在
	ErrCliUserExist = 40003
	// 客户端登陆信息错误
	ErrCliLogin = 40005

	// 服务端通用错误
	ErrorServer = 50000
)

var errMap = map[int]string{
	ErrorClient: "[Client Error] Undefined Error",
	ErrCliParam: "[Client Error] ",
	ErrCliUserExist: "[Client Error] User exist",
	ErrCliLogin:     "[Client Error] Username or Password error",

	ErrorServer: "[Server Error] Contact admin",
}

func GetErrMsg(code int) string {
	msg, err := errMap[code]
	if !err {
		return "[Undefined Error] Contact developer"
	}
	return msg
}
