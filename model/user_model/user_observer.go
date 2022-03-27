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
package user_model

import (
	"gorm.io/gorm"

	"gin-biz-web-api/pkg/hash"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	// 先判断用户密码是否已经经过 hash 加密
	if !hash.BcryptIsHashed(u.Password) {
		// 如果没有经过 hash 加密，则触发加密
		u.Password = hash.BcryptHash(u.Password)
	}

	return
}
