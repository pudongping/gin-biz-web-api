// 处理请求数据和表单验证
package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thedevsaddam/govalidator"

	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/responses"
)

// ValidateFunc 验证器中的验证方法
type ValidateFunc func(interface{}, *gin.Context) map[string][]string

// BindAndValidate 绑定并调用验证器方法验证参数
func BindAndValidate(c *gin.Context, obj interface{}, handler ValidateFunc) bool {
	response := responses.New(c)

	// 解析请求，支持 JSON 数据、表单请求和 URL Query
	// 参见：[gin 框架中文文档](https://www.kancloud.cn/shuangdeyu/gin_book/949426)
	// ShouldBindBodyWith https://github.com/gin-gonic/gin/pull/1341
	if err := c.ShouldBind(obj); err != nil {
		response.ToErrorResponse(errcode.BadRequest.WithDetails(err.Error()).WithError(errors.WithStack(err)), "请求体解析错误，请确认请求格式是否正确")
		return false
	}

	// 调用表单验证
	errs := handler(obj, c)

	// 判断验证是否通过
	if len(errs) > 0 {
		response.ToErrorValidateResponse(errcode.UnprocessableEntity, errs)
		return false
	}

	return true
}

// Validate 验证 form-data, x-www-form-urlencoded 和 query 传参类型的参数
// 如果要验证文件时 ref： https://github.com/thedevsaddam/govalidator/blob/master/doc/FILE_VALIDATION.md
func Validate(c *gin.Context, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置初始化
	opts := govalidator.Options{
		Request:  c.Request, // 请求实例对象
		Rules:    rules,     // 验证规则
		Messages: messages,  // 自定义错误消息
		// RequiredDefault: true,      // 所有的字段都要通过验证规则
		TagIdentifier: "valid", // 结构体中标签标识符
	}

	// 开始验证
	return govalidator.New(opts).Validate()
}

// ValidateJSON 验证 application/json 或者 text/plain 请求体数据，并绑定到指定结构上（比如：结构体、map）
// 绑定到结构体上 document link：https://github.com/thedevsaddam/govalidator/blob/master/doc/SIMPLE_STRUCT_VALIDATION.md
// 绑定到 map 上 document link：https://github.com/thedevsaddam/govalidator/blob/master/doc/MAP_VALIDATION.md
func ValidateJSON(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 结构体中标签标识符
	}

	return govalidator.New(opts).ValidateJSON()
}

// ValidateStruct 验证已有的结构体数据（不支持验证含有指针变量的结构体，但是支持验证 map[string]interface{}）
// document link： https://github.com/thedevsaddam/govalidator/blob/master/doc/STRUCT_VALIDATION.md
func ValidateStruct(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 结构体中标签标识符
	}

	return govalidator.New(opts).ValidateStruct()
}
