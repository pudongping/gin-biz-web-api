# 定时计划任务

> 使用了依赖包 [robfig/cron](https://github.com/robfig/cron) 基于此包做了简单的一些调整

## 使用示例

- 定义定时任务

在 `crontab` 目录下创建新文件，并定义 `Run()` 方法，例如：

`crontab/foo_crontab.go`

```go

package crontab

import (
	"gin-biz-web-api/pkg/logger"
)

type FooCrontab struct {
}

func (f FooCrontab) Run() {
	logger.Info("开始执行定时任务 FooCrontab")
}

```

- 设置执行定时任务

在 `bootstrap/crontab.go` 文件 `addScheduleTask()` 方法中添加计划任务，例如：

```go

package bootstrap

import (
	"fmt"

	crontabTask "gin-biz-web-api/crontab"
	"gin-biz-web-api/global"
)

func addScheduleTask() {

	// 每 5s 执行一次，也支持 Linux 的 crontab 命令的规则，具体请看
	// [robfig/cron](https://github.com/robfig/cron)  包文档
	fooCrontabEntryID, err := global.Crontab.AddJob("@every 5s", crontabTask.FooCrontab{})
	if err != nil {
		fmt.Println("错误 ====> ", err)
	}
	fmt.Println("fooCrontabEntryID ====> ", fooCrontabEntryID)

}

```