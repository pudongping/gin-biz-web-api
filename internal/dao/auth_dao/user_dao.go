package auth_dao

import (
	"gin-biz-web-api/model"
	"gin-biz-web-api/pkg/database"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (d *UserDao) GetUsers() (users []model.User, count int64) {
	database.DB.Where("id >= ?", 0).Find(&users).Count(&count)
	return
}
