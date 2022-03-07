package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	// 错误码
	code int
	// 错误提示信息
	msg string
	// 详细信息
	details []string
}

// 用于保存所有的错误码和错误提示信息
var codes = map[int]string{}

// NewError 添加一个错误码
func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

// Code 返回错误码
func (e *Error) Code() int {
	return e.code
}

// Msg 返回错误提示信息
func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args ...interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// Details 返回错误详细信息
func (e *Error) Details() []string {
	return e.details
}

// WithDetails 用以追加错误详情信息
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}

	return &newError
}

// Error 打印出错误码和错误提示信息
func (e *Error) Error() string {
	return fmt.Sprintf("错误码为：%d ，错误信息为： %s", e.Code(), e.Msg())
}

// HttpStatusCode 自身系统的状态码和 http 状态码的映射关系
func (e *Error) HttpStatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
