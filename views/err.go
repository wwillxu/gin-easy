package views

const (
	Success = 2000

	ErrParam        = 4001
	ErrAuth         = 4030
	ErrNilID        = 4002
	ErrInvalidID    = 4003
	ErrLogin        = 4005
	ErrUserNotExist = 4010
	ErrUserExist    = 4011
	ErrServer       = 5000
)

var errMap = map[int]string{
	Success:         "",
	ErrParam:        "[param error] bad request",
	ErrNilID:        "[param error] nil id",
	ErrInvalidID:    "[param error] invalid id",
	ErrLogin:        "[client error] wrong username or password",
	ErrUserNotExist: "[client error] user not exist",
	ErrUserExist:    "[client error] user already exist",
	ErrServer:       "[server error] contact admin",
}

func getErrMsg(code int) string {
	msg, err := errMap[code]
	if !err {
		return "[unknown error] contact admin"
	}
	return msg
}
