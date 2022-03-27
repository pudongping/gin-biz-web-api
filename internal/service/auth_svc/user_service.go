package auth_svc

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/internal/dao/auth_dao"
	"gin-biz-web-api/model/user_model"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (svc *UserService) GetUsers(c *gin.Context) ([]user_model.User, int64) {
	return auth_dao.NewUserDao().GetUsers()
}
