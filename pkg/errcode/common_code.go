// 公共的错误码
package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000, "服务器内部错误")
	InvalidParams             = NewError(10001, "入参错误")
	NotFound                  = NewError(10002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10007, "请求过多")
)
