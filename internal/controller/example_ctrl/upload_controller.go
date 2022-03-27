package example_ctrl

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/responses"
	"gin-biz-web-api/pkg/upload"
)

type UploadController struct {
}

// UploadFile 上传文件
// curl --location --request POST 'localhost:3000/api/example/upload-file' \
// --form 'file=@"/Users/pudongping/Downloads/12.jpg"' \
// --form 'type="image"'
func (ctrl *UploadController) UploadFile(c *gin.Context) {
	response := responses.New(c)

	file, fileHeader, err := c.Request.FormFile("file") // 读取入参 file 字段的上传文件信息
	if err != nil {
		response.ToErrorResponse(errcode.UnprocessableEntity.WithDetails(err.Error()).Msgf("文件"))
		return
	}

	fileType := c.PostForm("type") // 获取文件上传的类型
	if fileHeader == nil || fileType == "" {
		response.ToErrorResponse(errcode.UnprocessableEntity, "文件读取失败或上传类型不能为空")
		return
	}

	// 上传并保存文件
	fileInfo, err := upload.SaveUploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		response.ToErrorResponse(errcode.UnprocessableEntity.WithDetails(err.Error()), "文件上传失败")
		return
	}

	response.ToResponse(fileInfo)
}

// UploadAvatar 上传用户头像
// curl --location --request POST 'localhost:3000/api/example/upload-avatar' \
// --form 'avatar=@"/Users/pudongping/Downloads/11.jpg"'
func (ctrl *UploadController) UploadAvatar(c *gin.Context) {
	response := responses.New(c)

	file, fileHeader, err := c.Request.FormFile("avatar") // 读取入参 avatar 字段的上传文件信息
	if err != nil {
		response.ToErrorResponse(errcode.UnprocessableEntity.WithDetails(err.Error()).Msgf("头像文件"))
		return
	}

	fileInfo, err := upload.SaveUploadAvatar(file, fileHeader)

	if err != nil {
		response.ToErrorResponse(errcode.UnprocessableEntity.WithDetails(err.Error()), "文件上传失败")
		return
	}

	response.ToResponse(fileInfo)
}
