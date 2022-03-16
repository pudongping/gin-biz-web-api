// 定义验证码接口
// 受  base64Captcha.Store 的启发，可以参考他的代码设计流程
package verifycode

type Driver interface {

	// 保存验证码
	Set(key, value string) bool

	// 获取验证码
	Get(key string, isClear bool) string

	// 检查验证码
	Verify(key, answer string, isClear bool) bool
}
