package auth_ctrl

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/model/user_model"
	"gin-biz-web-api/pkg/auth"
	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/responses"
)

type UserController struct {
}

// Index 用户列表
// curl --location --request GET '0.0.0.0:3000/api/auth/user?page=3&per_page=2&order_by=id,desc|created_at,asc'
func (ctrl *UserController) Index(c *gin.Context) {

	response := responses.New(c)

	// var users []user_model.User
	// query := database.DB.Model(user_model.User{}).Where("id >= ?", 3)
	// paginate := paginator.Paginate(c, query, &users, 3)

	users, paginate := user_model.Paginate(c)

	if len(users) == 0 {
		response.ToErrorResponse(errcode.NotFound.Msgf("用户"))
		return
	}

	response.ToResponse(gin.H{
		"users":    users,
		"paginate": paginate,
	})

}

// Profile 用户个人信息
// curl --location --request GET '0.0.0.0:3000/api/auth/me' \
// --header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjQiLCJleHBpcmVfdGltZSI6MTY1MzEzMDM0MCwiZXhwIjoxNjUzMTMwMzQwLCJpYXQiOjE2NDc5NDYzNDAsImlzcyI6Imdpbi1iaXotd2ViLWFwaSIsIm5iZiI6MTY0Nzk0NjM0MH0.3Rzl8PmE519qWVmNziJ6ovH6Bwq5hnqmelkMUxfYsXc'
func (ctrl *UserController) Profile(c *gin.Context) {
	profile := auth.CurrentUser(c)
	responses.New(c).ToResponse(profile)
}
