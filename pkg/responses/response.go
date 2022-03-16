// 数据返回封装
package responses

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/logger"
)

type Response struct {
	Ctx *gin.Context
}

// New 实例化返回类
func New(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// ToResponse 正确数据返回
func (r *Response) ToResponse(data interface{}) {
	code := errcode.Success
	response := gin.H{"code": code.Code(), "msg": code.Msg()}
	if data == nil {
		response["data"] = gin.H{}
	} else {
		response["data"] = data
	}

	r.Ctx.JSON(http.StatusOK, response)
}

// ToErrorResponse 错误返回
func (r *Response) ToErrorResponse(err *errcode.Error, messages ...string) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}

	if len(messages) > 0 {
		response["msg"] = messages[0]
	}

	details := err.Details()
	if len(details) > 0 {
		if r.isShowDetails() {
			response["details"] = details
		} else {
			logger.ErrorJSON("ToErrorResponse", "details", details)
		}
	}

	r.Ctx.JSON(err.HttpStatusCode(), response)
}

// ToErrorValidateResponse 验证器验证不通过时，错误返回
// 返回的 json 示例为：
// {
//    "code": 100422,
//    "errors": {
//        "account": [
//            "账号为必填项",
//            "账号格式错误，只允许数字和英文",
//            "账号长度需在 3~20 之间"
//        ],
//        "email": [
//            "Email 为必填项",
//            "Email 长度需大于 4",
//            "Email 格式不正确，请提供有效的邮箱地址"
//        ],
//        "password": [
//            "密码为必填项",
//            "密码长度需大于 6"
//        ],
//        "password_confirm": [
//            "确认密码框为必填项"
//        ],
//        "verify_code": [
//            "验证码为必填",
//            "验证码长度必须为 6 位的数字"
//        ]
//    },
//    "msg": "验证码为必填"
// }
func (r *Response) ToErrorValidateResponse(err *errcode.Error, errors map[string][]string) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}

	if len(errors) > 0 {

		var kSlice []string
		for k := range errors {
			kSlice = append(kSlice, k)
		}

		sort.Strings(kSlice)

		for _, kk := range kSlice {
			response["msg"] = errors[kk][0]
			break
		}

		if r.isShowDetails() {
			response["details"] = errors
		}

	}

	r.Ctx.JSON(err.HttpStatusCode(), response)
}

// isShowDetails 本地环境或者开发环境且开启了 debug 模式，则在返回结果中显示错误详情信息
func (r *Response) isShowDetails() bool {
	return app.IsDebug() && (app.IsLocal() || app.IsDev())
}
