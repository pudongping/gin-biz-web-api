package auth_dao

import (
	"gin-biz-web-api/model/user_model"
	"gin-biz-web-api/pkg/database"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (d *UserDao) GetUsers() (users []user_model.User, count int64) {
	database.DB.Where("id >= ?", 3).Find(&users).Count(&count)
	return
}
