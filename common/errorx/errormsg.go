package errorx

var errorMsg map[int]string

const (
	ServerErrorCode              = 1000
	ParamErrorCode               = 1001
	CaptchaErrorCode             = 1002
	AccountErrorCode             = 1003
	PasswordErrorCode            = 1004
	NotPermMenuErrorCode         = 1005
	DeletePermMenuErrorCode      = 1006
	ParentPermMenuErrorCode      = 1007
	AddRoleErrorCode             = 1008
	DeleteRoleErrorCode          = 1009
	AddDeptErrorCode             = 1010
	DeleteDeptErrorCode          = 1011
	AddJobErrorCode              = 1012
	DeleteJobErrorCode           = 1013
	AddProfessionErrorCode       = 1014
	DeleteProfessionErrorCode    = 1015
	AddUserErrorCode             = 1016
	DeptHasUserErrorCode         = 1017
	RoleIsUsingErrorCode         = 1018
	ParentRoleErrorCode          = 1019
	ParentDeptErrorCode          = 1020
	AccountDisableErrorCode      = 1021
	SetParentIdErrorCode         = 1022
	SetParentTypeErrorCode       = 1023
	AddConfigErrorCode           = 1024
	AddDictionaryErrorCode       = 1025
	AuthErrorCode                = 1026
	DeleteDictionaryErrorCode    = 1027
	JobIsUsingErrorCode          = 1028
	ProfessionIsUsingErrorCode   = 1029
	ForbiddenErrorCode           = 1030
	UpdateRoleUniqueKeyErrorCode = 1031
	UpdateDeptUniqueKeyErrorCode = 1032
	AssigningRolesErrorCode      = 1033
	DeptIdErrorCode              = 1034
	ProfessionIdErrorCode        = 1035
	JobIdErrorCode               = 1036
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
	errorMsg[DeptHasUserErrorCode] = "该部门正在使用中"
	errorMsg[RoleIsUsingErrorCode] = "该角色正在使用中"
	errorMsg[ParentRoleErrorCode] = "父级角色不能为自己"
	errorMsg[ParentDeptErrorCode] = "父级部门不能为自己"
	errorMsg[AccountDisableErrorCode] = "账号已禁用"
	errorMsg[SetParentIdErrorCode] = "不能设置子级为自己的父级"
	errorMsg[SetParentTypeErrorCode] = "权限类型不能作为父级菜单"
	errorMsg[AddConfigErrorCode] = "配置已存在"
	errorMsg[AddDictionaryErrorCode] = "字典已存在"
	errorMsg[AuthErrorCode] = "授权已失效，请重新登录"
	errorMsg[DeleteDictionaryErrorCode] = "该字典集存在配置项"
	errorMsg[JobIsUsingErrorCode] = "该岗位正在使用中"
	errorMsg[ProfessionIsUsingErrorCode] = "该职称正在使用中"
	errorMsg[ForbiddenErrorCode] = "禁止操作"
	errorMsg[UpdateRoleUniqueKeyErrorCode] = "角色标识已存在"
	errorMsg[UpdateDeptUniqueKeyErrorCode] = "部门标识已存在"
	errorMsg[AssigningRolesErrorCode] = "角色不在可控范围"
	errorMsg[DeptIdErrorCode] = "部门不存在"
	errorMsg[ProfessionIdErrorCode] = "职称不存在"
	errorMsg[JobIdErrorCode] = "岗位不存在"
}

func MapErrMsg(errCode int) string {
	if msg, ok := errorMsg[errCode]; ok {
		return msg
	} else {
		return "服务繁忙，请稍后重试"
	}
}
