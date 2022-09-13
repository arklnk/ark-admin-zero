package errorx

import "ark-admin-zero/config"

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(code int) error {
	return NewCodeError(code, MapErrMsg(code))
}

func NewHandlerError(code int, msg string) error {
	return NewCodeError(code, msg)
}

func NewSystemError(code int, msg string) error {
	if config.SysShowSystemError {
		return NewCodeError(code, msg)
	} else {
		return NewCodeError(code, MapErrMsg(code))
	}
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
