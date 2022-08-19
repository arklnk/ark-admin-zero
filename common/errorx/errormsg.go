package errorx

var errorMsg map[int]string

const (
	ServerErrorCode               = 1000
	ParamErrorCode                = 1001
	CaptchaErrorCode              = 1002
	AccountErrorCode              = 1003
	PasswordErrorCode             = 1004
	NotPermMenuErrorCode          = 1005
	DeletePermMenuErrorCode       = 1006
	ParentPermMenuErrorCode       = 1007
	AddRoleErrorCode              = 1008
	DeleteRoleErrorCode           = 1009
	AddDeptErrorCode              = 1010
	DeleteDeptErrorCode           = 1011
	AddJobErrorCode               = 1012
	DeleteJobErrorCode            = 1013
	AddProfessionErrorCode        = 1014
	DeleteProfessionErrorCode     = 1015
	AddUserErrorCode              = 1016
	DeleteUserErrorCode           = 1017
	DeptHasUserErrorCode          = 1018
	RoleIsUsingErrorCode          = 1019
	ParentRoleErrorCode           = 1020
	ParentDeptErrorCode           = 1021
)

func init() {
	errorMsg = make(map[int]string)
	errorMsg[ServerErrorCode] = "服务繁忙，请稍后重试"
	errorMsg[CaptchaErrorCode] = "验证码错误"
	errorMsg[AccountErrorCode] = "账号错误"
	errorMsg[PasswordErrorCode] = "密码错误"
	errorMsg[NotPermMenuErrorCode] = "权限不足"
	errorMsg[DeletePermMenuErrorCode] = "该权限菜单存在子级权限菜单"
	errorMsg[ParentPermMenuErrorCode] = "父级菜单不能为自己"
	errorMsg[AddRoleErrorCode] = "角色已存在"
	errorMsg[DeleteRoleErrorCode] = "该角色存在子角色"
	errorMsg[AddDeptErrorCode] = "部门已存在"
	errorMsg[DeleteDeptErrorCode] = "该部门存在子部门"
	errorMsg[AddJobErrorCode] = "岗位已存在"
	errorMsg[DeleteJobErrorCode] = "该岗位正在使用中"
	errorMsg[AddProfessionErrorCode] = "职称已存在"
	errorMsg[DeleteProfessionErrorCode] = "该职称正在使用中"
	errorMsg[AddUserErrorCode] = "账号已存在"
	errorMsg[DeleteUserErrorCode] = "禁止刪除超级管理员"
	errorMsg[DeptHasUserErrorCode] = "该部门正在使用中"
	errorMsg[RoleIsUsingErrorCode] = "该角色正在使用中"
	errorMsg[ParentRoleErrorCode] = "父级角色不能为自己"
	errorMsg[ParentDeptErrorCode] = "父级部门不能为自己"

}

func MapErrMsg(errCode int) string {
	if msg, ok := errorMsg[errCode]; ok {
		return msg
	} else {
		return "服务繁忙，请稍后重试"
	}
}
