package errorx

var ErrorMsg map[int]string

const (
	CaptchaErrorCode = 1002
	ParamErrorCode   = 1003
)

func init() {
	ErrorMsg = make(map[int]string)
	ErrorMsg[CaptchaErrorCode] = "验证码错误"
}

func MapErrMsg(errCode int) string {
	if msg, ok := ErrorMsg[errCode]; ok {
		return msg
	} else {
		return "错误码未定义"
	}
}
