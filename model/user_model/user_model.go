package user_model

import (
	"gin-biz-web-api/model"
)

type User struct {
	*model.BaseModel

	// 账号  varchar(255) is_nullable NO
	Account string `json:"account"`
	// 邮箱  varchar(80) is_nullable YES
	Email string `json:"email"`
	// 手机号  varchar(40) is_nullable YES
	Phone string `json:"phone"`
	// 密码  varchar(255) is_nullable NO
	Password string `json:"password"`
	// 省份  varchar(20) is_nullable NO
	Province string `json:"province"`
	// 市区  varchar(40) is_nullable NO
	City string `json:"city"`
	// 区县  varchar(255) is_nullable NO
	Country string `json:"country"`
	// 昵称  varchar(255) is_nullable NO
	Nickname string `json:"nickname"`
	// 自我简介  text is_nullable YES
	Introduction string `json:"introduction"`
	// 头像地址  varchar(255) is_nullable NO
	Avatar string `json:"avatar"`

	*model.CommonTimestampsField
}

func (u User) TableName() string {
	return "users"
}
