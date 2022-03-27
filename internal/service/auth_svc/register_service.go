package auth_svc

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/internal/requests/auth_request"
	"gin-biz-web-api/model/user_model"
	"gin-biz-web-api/pkg/database"
	"gin-biz-web-api/pkg/jwt"
)

type RegisterService struct {
}

func NewRegisterService() *RegisterService {
	return &RegisterService{}
}

// CreateUserToken 创建用户并返回 token
func (svc *RegisterService) CreateUserToken(c *gin.Context, request auth_request.SignupUsingEmailRequest) string {

	user := user_model.User{
		Account:  request.Account,
		Email:    request.Email,
		Password: request.Password,
	}

	database.DB.Model(&user_model.User{}).Select("account", "email", "password").Create(&user)

	if user.ID > 0 {
		// 生成 token
		return jwt.NewJWT().GenerateToken(user.GetStringID())
	} else {
		return ""
	}

}
