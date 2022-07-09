package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"gin-biz-web-api/global"

	"go.uber.org/zap"
)

// Rotate 按日期轮转日志文件
// maxSize 每个日志文件保存的最大尺寸，单位 M
// maxFile 最多保存多少天前的日志
func Rotate(maxSize, maxFile int64) error {
	var err error
	rootPath := global.RootPath
	logPath := path.Join(rootPath, LogPath)

	// 每个日志文件保存的最大尺寸 单位：M
	maxSize = maxSize * 1024 * 1024

	dayBefore := time.Now().AddDate(0, 0, -1*int(maxFile))
	dayBeforeUnixTime := time.Date(
		dayBefore.Year(),
		dayBefore.Month(),
		dayBefore.Day(),
		0,
		0,
		0,
		0,
		dayBefore.Location(),
	).Unix()

	logFiles, err := ioutil.ReadDir(logPath)
	if err != nil {
		return err
	}

	logFiles = sortByTime(logFiles)

	var needRemoveFL []os.FileInfo

	for _, fileInfo := range logFiles {
		if fileInfo.IsDir() || !strings.HasSuffix(fileInfo.Name(), ".log") {
			continue
		}

		if fileInfo.ModTime().Unix() <= dayBeforeUnixTime {
			// 日志文件修改时间小于最多保存天数前的时间时，需要被删除
			needRemoveFL = append(needRemoveFL, fileInfo)
		} else {
			if fileInfo.Size() >= maxSize {
				// 或者文件在需要保存天数内，但是文件大小超过限制时，也需要被删除
				needRemoveFL = append(needRemoveFL, fileInfo)
			}
		}
	}

	for _, fileInfo := range needRemoveFL {
		Info("开始自动清理日志文件",
			zap.String("文件名称", fileInfo.Name()),
			zap.String("文件大小", fmt.Sprintf("%.2fMB", float64(fileInfo.Size())/float64(1024*1024))),
			zap.Time("文件 mtime", fileInfo.ModTime()),
		)
		absLogFile := logPath + "/" + fileInfo.Name()
		err := os.Remove(absLogFile)
		if err != nil {
			Error("自动清理日志文件失败",
				zap.String("文件绝对路径", absLogFile),
				zap.Error(err),
			)
		} else {
			Info("自动清理日志文件成功", zap.String("文件绝对路径", absLogFile))
		}
	}

	return nil
}

func sortByTime(fl []os.FileInfo) []os.FileInfo {
	sort.Slice(fl, func(i, j int) bool {
		flag := false
		if fl[i].ModTime().After(fl[j].ModTime()) {
			flag = true
		} else if fl[i].ModTime().Equal(fl[j].ModTime()) {
			if fl[i].Name() < fl[j].Name() {
				flag = true
			}
		}
		return flag
	})
	return fl
}
