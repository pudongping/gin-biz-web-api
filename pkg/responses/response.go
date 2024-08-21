// 数据返回封装
package responses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/helper/mapx"
	"gin-biz-web-api/pkg/logger"
	"gin-biz-web-api/pkg/paginator"
)

type Response struct {
	Ctx *gin.Context
}

type ResponseData struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 错误信息
	Data interface{} `json:"data"` // 数据
}

// New 实例化返回类
func New(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// ToResponse 正确数据返回
func (r *Response) ToResponse(data interface{}) {
	code := errcode.Success
	response := ResponseData{
		Code: code.Code(),
		Msg:  code.Msg(),
		Data: nil,
	}
	if data == nil {
		response.Data = gin.H{}
	} else {
		response.Data = data
	}

	r.Ctx.JSON(http.StatusOK, response)
}

// ToResponseWithPagination 返回分页数据
func (r *Response) ToResponseWithPagination(result interface{}, pagination paginator.Pagination) {
	r.ToResponse(gin.H{
		"result":     result,
		"pagination": pagination,
	})
}

// ToErrorResponse 错误返回
func (r *Response) ToErrorResponse(err error, messages ...string) {
	var e = errcode.InternalServerError.WithError(err).WithDetails(err.Error())

	if customErr, ok := err.(*errcode.Error); ok {
		e = customErr
	}

	response := ResponseData{
		Code: e.Code(),
		Msg:  e.Msg(),
		Data: nil,
	}

	if len(messages) > 0 {
		response.Msg = messages[0]
	}

	details := e.Details()
	if len(details) > 0 && r.isShowDetails() {
		logger.ErrorJSON("ToErrorResponse", "details", details)
	}

	if r.isShowDetails() && e.Err() != nil {
		fmt.Printf("原始错误详情信息 ====> %+v \n", e.Err())
	}

	// r.Ctx.JSON(err.HttpStatusCode(), response)
	// r.Ctx.AbortWithStatusJSON(err.HttpStatusCode(), response)
	r.Ctx.AbortWithStatusJSON(http.StatusOK, response)
}

// ToErrorValidateResponse 验证器验证不通过时，错误返回
func (r *Response) ToErrorValidateResponse(err *errcode.Error, errors map[string][]string) {
	response := ResponseData{
		Code: err.Code(),
		Msg:  err.Msg(),
		Data: nil,
	}

	if len(errors) > 0 {
		ks := mapx.SortAscKeyString(errors)

		for _, k := range ks {
			response.Msg = errors[k][0]
			break
		}

		if r.isShowDetails() {
			logger.ErrorJSON("ToErrorResponse", "details", errors)
		}
	}

	// r.Ctx.AbortWithStatusJSON(err.HttpStatusCode(), response)
	r.Ctx.AbortWithStatusJSON(http.StatusOK, response)
}

// isShowDetails 本地环境或者开发环境且开启了 debug 模式，则在返回结果中显示错误详情信息
func (r *Response) isShowDetails() bool {
	return app.IsDebug() && (app.IsLocal() || app.IsDev())
}
