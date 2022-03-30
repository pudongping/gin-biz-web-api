package controller

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/responses"
)

type TestController struct {
}

// curl --location --request GET '0.0.0.0:3000/api/test'
func (ctrl *TestController) Test(c *gin.Context) {
	response := responses.New(c)
	response.ToResponse(nil)
}

// curl --location --request GET '0.0.0.0:3000/api/test/tt'
func (ctrl *TestController) Tt(c *gin.Context) {
	response := responses.New(c)
	response.ToResponse(nil)
}
