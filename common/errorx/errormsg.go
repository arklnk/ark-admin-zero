package errorx

var errorMsg map[int]string

const (
	ServerErrorCode           = 1000
	ParamErrorCode            = 1001
	CaptchaErrorCode          = 1002
	AccountErrorCode          = 1003
	PasswordErrorCode         = 1004
	NotPermMenuErrorCode      = 1005
	DeletePermMenuErrorCode   = 1006
	AddRoleErrorCode          = 1007
	DeleteRoleErrorCode       = 1008
	AddDeptErrorCode          = 1008
	DeleteDeptErrorCode       = 1010
	AddJobErrorCode           = 1011
	DeleteJobErrorCode        = 1012
	AddProfessionErrorCode    = 1013
	DeleteProfessionErrorCode = 1014
	AddUserErrorCode          = 1015
)

func init() {
	errorMsg = make(map[int]string)
	errorMsg[ServerErrorCode] = "服务繁忙，请稍后重试"
	errorMsg[CaptchaErrorCode] = "验证码错误"
	errorMsg[AccountErrorCode] = "账号错误"
	errorMsg[PasswordErrorCode] = "密码错误"
	errorMsg[NotPermMenuErrorCode] = "权限不足"
	errorMsg[DeletePermMenuErrorCode] = "该权限菜单存在子级权限菜单"
	errorMsg[AddRoleErrorCode] = "角色已存在"
	errorMsg[DeleteRoleErrorCode] = "该角色存在子角色"
	errorMsg[AddDeptErrorCode] = "部门已存在"
	errorMsg[DeleteDeptErrorCode] = "该部门存在子部门"
	errorMsg[AddJobErrorCode] = "岗位已存在"
	errorMsg[DeleteJobErrorCode] = "该岗位正在使用中"
	errorMsg[AddProfessionErrorCode] = "职称已存在"
	errorMsg[DeleteProfessionErrorCode] = "该职称正在使用中"
	errorMsg[AddUserErrorCode] = "账号已存在"
}

func MapErrMsg(errCode int) string {
	if msg, ok := errorMsg[errCode]; ok {
		return msg
	} else {
		return "服务繁忙，请稍后重试"
	}
}
