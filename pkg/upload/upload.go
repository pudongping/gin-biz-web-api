// 上传
package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"

	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/file"
	"gin-biz-web-api/pkg/helper/strx"
	"gin-biz-web-api/pkg/logger"
)

type FileType string

const (
	TypeAvatar FileType = "avatar" // 用户头像
	TypeImage  FileType = "image"  // 图片
	TypeFile   FileType = "file"   // 文件
)

type FileInfo struct {
	OriginFileName string `json:"origin_file_name"` // 原始文件名
	FileName       string `json:"file_name"`        // 文件名
	AbsPath        string `json:"abs_path"`         // 文件绝对路径（相对本项目的绝对路径）
	AccessUrl      string `json:"access_url"`       // 文件资源访问地址
}

// SaveUploadFile 保存上传文件
func SaveUploadFile(fileType FileType, files multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	// 重新生成文件名 eg：`c20ad4d76fe97759aa27a0c99bff6710-20220319015124.jpg`
	fileName := file.GenNewFileName(fileHeader.Filename)

	// 检查文件后缀是否包含在约定的后缀配置项中
	if !CheckContainExt(fileType, fileName) {
		return nil, errors.New("文件后缀不支持")
	}

	// 检查文件大小是否超出最大大小限制
	if CheckMaxSize(fileType, files) {
		return nil, errors.New("超过最大文件限制")
	}

	publicPath := config.GetString("upload.save_path") // `public/uploads`

	dirName := fmt.Sprintf("%s/%s", fileType, app.TimeNowInTimezone().Format("2006/01/02")) // `image/2022/03/19`

	// 上传路径 eg：`public/uploads/image/2022/03/19`
	uploadSavePath := fmt.Sprintf("%s/%s", publicPath, dirName)

	// 检查目录是否存在
	if _, isExists := file.IsExists(uploadSavePath); !isExists {
		// 创建文件夹
		if err := file.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("无法创建保存目录：" + uploadSavePath)
		}
	}

	// 检查是否有写入权限
	if file.CheckPermission(uploadSavePath) {
		return nil, errors.New("权限不足")
	}

	// public/uploads/image/2022/03/19/c20ad4d76fe97759aa27a0c99bff6710-20220319015124.jpg
	dst := filepath.Join(uploadSavePath, fileName)
	// 保存文件
	if err := file.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	// http://localhost:3000/static/image/2022/03/19/c20ad4d76fe97759aa27a0c99bff6710-20220319023344.jpg
	accessUrl := fmt.Sprintf(
		"%s/%s/%s",
		app.URL(config.GetString("upload.static_fs_relative_path")), // `http://localhost:3000/static`
		dirName,                                                     // `image/2022/03/19`
		fileName,                                                    // `c20ad4d76fe97759aa27a0c99bff6710-20220319023344.jpg`
	)

	return &FileInfo{
		OriginFileName: fileHeader.Filename,
		FileName:       fileName,
		AbsPath:        dst,
		AccessUrl:      accessUrl,
	}, nil

}

// SaveUploadAvatar 上传头像
func SaveUploadAvatar(files multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {

	// 先按照正常的流程上传头像图片
	fileInfo, err := SaveUploadFile(TypeAvatar, files, fileHeader)
	if err != nil {
		return nil, err
	}

	// 对头像进行裁剪
	resizeFileInfo, err := TailoringImage(TypeAvatar, fileHeader, fileInfo.AbsPath, 256, 256)
	if err != nil {
		logger.LogErrorIf(err)
		return fileInfo, nil
	}

	// 删除旧的头像
	err = os.Remove(fileInfo.AbsPath)
	if err != nil {
		logger.LogErrorIf(err)
		return fileInfo, nil
	}

	return resizeFileInfo, nil
}

// TailoringImage 裁剪图片
// fileType 文件类型
// fileHeader
// src 原图片路径，相对项目的绝对路径 eg：`public/uploads/image/2022/03/19/c20ad4d76fe97759aa27a0c99bff6710-20220319015124.jpg`
// width 所需要裁剪的宽度
// height 所需要裁剪的高度
func TailoringImage(fileType FileType, fileHeader *multipart.FileHeader, src string, width, height int) (*FileInfo, error) {

	var fileInfo FileInfo

	// 打开图片
	img, err := imaging.Open(src, imaging.AutoOrientation(true))
	if err != nil {
		return nil, err
	}

	// 采用高清模式裁剪成缩略图
	resize := imaging.Thumbnail(img, width, height, imaging.Lanczos)

	// 重新生成文件名 eg：`PqOem-c20ad4d76fe97759aa27a0c99bff6710-20220319015124.jpg`
	fileName := strx.StrRandomString(5) + "-" + file.GenNewFileName(fileHeader.Filename)
	publicPath := config.GetString("upload.save_path")                                      // `public/uploads`
	dirName := fmt.Sprintf("%s/%s", fileType, app.TimeNowInTimezone().Format("2006/01/02")) // `image/2022/03/19`

	// 保存路径 eg：`public/uploads/image/2022/03/19/PqOem-c20ad4d76fe97759aa27a0c99bff6710-20220319015124.jpg`
	filePath := fmt.Sprintf("%s/%s/%s", publicPath, dirName, fileName)

	// 保存文件到指定路径
	err = imaging.Save(resize, filePath)
	if err != nil {
		return nil, err
	}

	// http://localhost:3000/static/image/2022/03/19/PqOem-c20ad4d76fe97759aa27a0c99bff6710-20220319023344.jpg
	accessUrl := fmt.Sprintf(
		"%s/%s/%s",
		app.URL(config.GetString("upload.static_fs_relative_path")), // `http://localhost:3000/static`
		dirName,                                                     // `image/2022/03/19`
		fileName,                                                    // `PqOem-c20ad4d76fe97759aa27a0c99bff6710-20220319023344.jpg`
	)

	fileInfo.OriginFileName = fileHeader.Filename
	fileInfo.FileName = fileName
	fileInfo.AbsPath = filePath
	fileInfo.AccessUrl = accessUrl

	return &fileInfo, nil
}

// CheckContainExt 检查文件后缀是否包含在约定的后缀配置项中
// true 包含，false 不包含
func CheckContainExt(t FileType, fileName string) bool {
	ext := path.Ext(fileName) // 获取文件后缀名 eg：`.jpg`
	ext = strings.ToLower(ext)
	imageAllows := config.GetStringSlice("upload.image.allow_suffix")
	fileAllows := config.GetStringSlice("upload.file.allow_suffix")

	switch t {
	case TypeAvatar:
		fallthrough
	case TypeImage:
		for _, allowExt := range imageAllows {
			if strings.ToLower(allowExt) == ext {
				return true
			}
		}
	case TypeFile:
		for _, allowExt := range fileAllows {
			if strings.ToLower(allowExt) == ext || "*" == allowExt {
				return true
			}
		}
	}

	return false
}

// CheckMaxSize 检查文件大小是否超出了最大限制
// true 超出，false 没有超出
func CheckMaxSize(t FileType, f multipart.File) bool {

	imageMaxSize := config.GetInt("upload.image.max_size")
	fileMaxSize := config.GetInt("upload.file.max_size")
	maxSize := config.GetInt("upload.max_size") * 1024 * 1024

	switch t {
	case TypeAvatar:
		fallthrough
	case TypeImage:
		if imageMaxSize > 0 {
			maxSize = imageMaxSize * 1024 * 1024
		}
	case TypeFile:
		if fileMaxSize > 0 {
			maxSize = fileMaxSize * 1024 * 1024
		}
	}

	return file.CheckMaxSize(f, maxSize)
}
