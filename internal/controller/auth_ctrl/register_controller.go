package auth_ctrl

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/internal/requests/auth_request"
	"gin-biz-web-api/internal/service/auth_svc"
	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/responses"
	"gin-biz-web-api/pkg/validator"
)

type RegisterController struct {
}

// SignupUsingEmail 使用邮箱进行注册
// curl --location --request POST '0.0.0.0:3000/api/auth/register/using-email' \
// --header 'Content-Type: application/x-www-form-urlencoded' \
// --data-urlencode 'account=alex' \
// --data-urlencode 'email=123456@qq.com' \
// --data-urlencode 'password=123456' \
// --data-urlencode 'password_confirm=123456' \
// --data-urlencode 'verify_code=123456'
func (ctrl *RegisterController) SignupUsingEmail(c *gin.Context) {
	response := responses.New(c)

	// 表单验证
	request := auth_request.SignupUsingEmailRequest{}
	if ok := validator.BindAndValidate(c, &request, auth_request.SignupUsingEmail); !ok {
		return
	}

	token := auth_svc.NewRegisterService().CreateUserToken(c, request)

	if "" == token {
		response.ToErrorResponse(errcode.DBError)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})

}
