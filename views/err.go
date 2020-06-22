package views

const (
	Success = 20000

	ErrorClient     = 40000
	ErrCliParam     = 40003
	ErrCliInvalidID = 40005
	ErrCliUserExist = 40001
	ErrCliLogin     = 40002

	ErrorServer = 50000
)

var errMap = map[int]string{
	ErrorClient:     "[Client Error] Undefined Error",
	ErrCliParam:     "[Client Error] ",
	ErrCliInvalidID: "[Client Error] Pass invalid ID",
	ErrCliUserExist: "[Client Error] User exist",
	ErrCliLogin:     "[Client Error] Username or Password error",

	ErrorServer: "[Server Error] Undefined Error",
}

func GetErrMsg(code int) string {
	msg, err := errMap[code]
	if !err {
		return "[Undefined Error] Contact developer"
	}
	return msg
}
