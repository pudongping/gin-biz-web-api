// 这里专门用来写很简单的 sql 语句
package user_model

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/database"
	"gin-biz-web-api/pkg/paginator"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (u *User) Create() {
	database.DB.Create(&u)
}

// Save 更新用户信息，通过影响行数来判断是否更新成功
func (u *User) Save() int64 {
	return database.DB.Save(&u).RowsAffected
}

// Delete 删除用户，通过影响行数来判断是否删除成功
func (u *User) Delete() int64 {
	return database.DB.Delete(&u).RowsAffected
}

// All 获取所有的用户信息
func All() (users []User) {
	database.DB.Find(&users)
	return
}

// GetOne 通过 ID 获取用户
func GetOne(id string) (user User) {
	database.DB.First(&user, id)
	return
}

// Paginate 分页
func Paginate(c *gin.Context, perPage ...int) (users []User, paging paginator.Pagination) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		perPage...,
	)
	return
}
