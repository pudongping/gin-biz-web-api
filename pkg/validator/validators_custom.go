package validator

import (
	"gin-biz-web-api/pkg/captcha"
	"gin-biz-web-api/pkg/verifycode"
)

// ValidationPasswordConfirm 自定义规则，检查两次密码是否正确
func ValidationPasswordConfirm(password, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}

// ValidationCaptcha 自定义规则，验证【图片验证码】
func ValidationCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

// ValidationVerifyCode 自定义规则，验证【验证码】
func ValidationVerifyCode(key, answer string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerifyCode().CheckVerifyCode(key, answer); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}
	return errs
}
