package errorx

var errorMsg map[int]string

const (
	ParamErrorCode   = 1001
	CaptchaErrorCode = 1002
)

func init() {
	errorMsg = make(map[int]string)
	errorMsg[CaptchaErrorCode] = "验证码错误"
}

func MapErrMsg(errCode int) string {
	if msg, ok := errorMsg[errCode]; ok {
		return msg
	} else {
		return "服务繁忙，请稍后重试"
	}
}
