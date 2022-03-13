// 公共的错误码
package errcode

var (
	Success              = NewError(0, "请求成功")
	Fail                 = NewError(1, "请求失败")
	BadRequest           = NewError(100400, "请求异常")
	Unauthorized         = NewError(100401, "无权操作")
	Forbidden            = NewError(100403, "禁止操作")
	NotFound             = NewError(100404, "找不到数据")
	MethodNotAllowed     = NewError(100405, "不允许此请求方法")
	RequestTimeout       = NewError(100408, "请求超时")
	UnsupportedMediaType = NewError(100415, "请求体错误")
	UnprocessableEntity  = NewError(100422, "%s参数校验错误")
	TooManyRequests      = NewError(100429, "请求太频繁")
	InternalServerError  = NewError(100500, "服务器内部错误")
	BadGateway           = NewError(100502, "网关错误")
	GatewayTimeout       = NewError(100504, "网关超时")
	DBError              = NewError(100600, "数据库操作失败")
)
