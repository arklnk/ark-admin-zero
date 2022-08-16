package errorx

var errorMsg map[int]string

const (
	ServerErrorCode      = 1000
	ParamErrorCode       = 1001
	CaptchaErrorCode     = 1002
	AccountErrorCode     = 1003
	PasswordErrorCode    = 1004
	NotPermMenuErrorCode = 1005
	AddRoleErrorCode     = 1006
	DeleteRoleErrorCode  = 1007
)

func init() {
	errorMsg = make(map[int]string)
	errorMsg[ServerErrorCode] = "服务繁忙，请稍后重试"
	errorMsg[CaptchaErrorCode] = "验证码错误"
	errorMsg[AccountErrorCode] = "账号错误"
	errorMsg[PasswordErrorCode] = "密码错误"
	errorMsg[NotPermMenuErrorCode] = "权限不足"
	errorMsg[AddRoleErrorCode] = "角色已存在"
	errorMsg[DeleteRoleErrorCode] = "该角色存在子角色"
}

func MapErrMsg(errCode int) string {
	if msg, ok := errorMsg[errCode]; ok {
		return msg
	} else {
		return "服务繁忙，请稍后重试"
	}
}
