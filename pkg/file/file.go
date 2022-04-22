// package file 负责处理文件相关
package file

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"gin-biz-web-api/pkg/helper"
	"gin-biz-web-api/pkg/helper/strx"
)

// IsExists 检查目录或者文件是否存在
// true 存在，false 不存在
func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path) // 获取文件的描述信息 FileInfo
	return f, err == nil || os.IsExist(err)
}

// Put 将数据存入文件中
func Put(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

// NameWithoutExtension 返回不带文件后缀的文件名称
func NameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// GenNewFileName 生成新的文件名称 eg：c20ad4d76fe97759aa27a0c99bff6710-20220319015124.jpg
func GenNewFileName(filename string) string {
	ext := path.Ext(filename)              // 获取文件后缀
	name := NameWithoutExtension(filename) // 只有文件名，没有后缀
	// 将原始文件名加密成 md5 然后拼接当前时间和随机字符串作为新的文件名
	newName := strx.StrRandomOptionalString(8, strx.LowerCase+strx.UpperCase) + "-" + helper.EncodeMD5(name) + "-" + time.Now().Format("20060102150405")

	// 新的文件名称
	return newName + ext
}

// CheckMaxSize 检查文件大小是否超出最大限制（kb 进行比较）
// true 超出，false 没有超出
func CheckMaxSize(f multipart.File, maxSize int) bool {
	content, _ := ioutil.ReadAll(f)
	return len(content) >= maxSize
}

// CheckPermission 检查文件权限是否足够
// true 不够 false 足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// 创建在上传文件时所使用的保存目录
func CreateSavePath(dst string, perm os.FileMode) error {
	// 该方法将会以传入的 os.FileMode 权限位去递归创建所需的所有目录结构，
	// 若涉及的目录均已存在，则不会进行任何操作，直接返回 nil
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// SaveFile 保存文件
// 参考于：gin.Context.SaveUploadedFile() 方法
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open() // 打开源地址的文件
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst) // 创建目标地址的文件
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
