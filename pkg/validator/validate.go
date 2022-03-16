// 处理请求数据和表单验证
package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"

	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/responses"
)

// ValidateFunc 验证器的函数
type ValidateFunc func(interface{}, *gin.Context) map[string][]string

// BindAndValidate 控制器中调用
func BindAndValidate(c *gin.Context, obj interface{}, handler ValidateFunc) bool {
	response := responses.New(c)

	// 解析请求，支持 JSON 数据、表单请求和 URL Query
	// 参见：[gin 框架中文文档](https://www.kancloud.cn/shuangdeyu/gin_book/949426)
	if err := c.ShouldBind(obj); err != nil {
		response.ToErrorResponse(errcode.BadRequest.WithDetails(err.Error()), "请求体解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数建议使用 JSON 格式")
		return false
	}

	// 表单验证
	errs := handler(obj, c)

	// 判断验证是否通过
	if len(errs) > 0 {
		response.ToErrorValidateResponse(errcode.UnprocessableEntity, errs)
		return false
	}

	return true
}

// Validate 验证器内部使用，用于验证表单数据
func Validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 结构体中标签标识符
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}

// ValidateFile 验证器内部使用，用于验证文件
func ValidateFile(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 结构体中标签标识符
	}

	// 调用 govalidator 的 Validate 方法来验证文件
	return govalidator.New(opts).Validate()
}
