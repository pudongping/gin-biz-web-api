// 数据返回封装
package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

// NewResponse 实例化返回类
func NewResponse(ctx *gin.Context) *Response {
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
		response["details"] = details
	}

	r.Ctx.JSON(err.HttpStatusCode(), response)
}

// ToErrorValidateResponse 验证器验证不通过时，错误返回
func (r *Response) ToErrorValidateResponse(err *errcode.Error, errors map[string][]string) {
	response := gin.H{"code": err.Code(), "msg": err.Msg(), "errors": errors}

	if len(errors) > 0 {
		for k := range errors {
			response["msg"] = errors[k][0]
			break
		}
	}

	r.Ctx.JSON(err.HttpStatusCode(), response)
}
