package e

const (
	Success               = 200
	Error                 = 500
	ErrorInputMissing     = 40001
	ErrorInputType        = 40002
	ErrorUsernameExist    = 40003
	ErrorUsernameNotExist = 40004
	ErrorTokenInit        = 40010
	ErrorPassword         = 40005
	ErrorRegister         = 50001
	ErrorGetUID           = 50002
	ErrorUpdateUser       = 50003
)

var MsgMap = map[int]string{
	Error:                 "发生错误",
	ErrorInputMissing:     "缺少输入数据",
	ErrorInputType:        "错误输入类型",
	ErrorUsernameExist:    "用户名已存在",
	ErrorUsernameNotExist: "用户不存在",
	ErrorTokenInit:        "Token生成失败",
	ErrorPassword:         "密码错误",
	ErrorRegister:         "注册失败",
	ErrorGetUID:           "获取登录信息失败",
	ErrorUpdateUser:       "资料更新失败",
}

func GetMsg(code int) string {
	msg, err := MsgMap[code]
	if !err {
		return "get err failed"
	}
	return msg
}
