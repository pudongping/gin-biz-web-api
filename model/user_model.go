package model

import (
	"gorm.io/gorm"

	"gin-biz-web-api/pkg/hash"
)

type User struct {
	*BaseModel

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

	*CommonTimestampsField
}

func (u User) TableName() string {
	return "users"
}

// 模型钩子
// 参见： [gorm document link](https://gorm.io/zh_CN/docs/hooks.html)
// func (u *User) BeforeSave(tx *gorm.DB) (err error) {}
// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {}
// func (u *User) AfterCreate(tx *gorm.DB) (err error) {}
// func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {}
// func (u *User) AfterUpdate(tx *gorm.DB) (err error) {}
// func (u *User) AfterSave(tx *gorm.DB) (err error) {}
// func (u *User) BeforeDelete(tx *gorm.DB) (err error) {}
// func (u *User) AfterDelete(tx *gorm.DB) (err error) {}
// func (u *User) AfterFind(tx *gorm.DB) (err error) {}

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	// 先判断用户密码是否已经经过 hash 加密
	if !hash.BcryptIsHashed(u.Password) {
		// 如果没有经过 hash 加密，则触发加密
		u.Password = hash.BcryptHash(u.Password)
	}

	return
}
