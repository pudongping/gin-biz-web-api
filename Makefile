# 这一行的意义在于告诉 make 命令，以下的命令都是伪命令，不是真正的文件名，这样即使存在同名的文件， make 命令也不会因为找到文件而不执行命令，依然会去执行这个目标下的命令。
.PHONY: help all build build-linux build-windows build-darwin version clean gotool

AppName=gin-biz-web-api
# 打包时的软件版本号
BuildVersion=v1.0.0
# v1.2.0-11-g98d9b66-dev
Version := $(shell git describe --tags --always --dirty="-dev")
# 98d9b668dd382b6f0f110782667a3ebf882109ee
GitCommit := $(shell git rev-list -1 HEAD)
# 当前时间
CurrentTime := $(shell date +%Y-%m-%d,%H:%M:%S)
# 构建二进制文件时 ldflags 所需参数
LdflagsArgs := "-X main.buildTime=$(CurrentTime) -X main.buildVersion=$(BuildVersion) -X 'main.goVersion=$(shell go version)' -X main.gitCommitID=$(shell git rev-parse HEAD)"

# 直接运行 make 命令时，会执行第一个目标，这里是 help，也就是说直接运行 make 命令时，相当于执行 make help 命令。
help:
	@echo "Usage: make <option>"
	@echo "\033[1;31;44mOptions and effects:\033[0m"
	@echo "  help           : Show help."
	@echo "  all            : Build multiple binary of this project."
	@echo "  build          : Build the binary of this project for current platform."
	@echo "  build-linux    : Build the linux binary of this project."
	@echo "  build-windows  : Build the windows binary of this project."
	@echo "  build-darwin   : Build the darwin binary of this project."
	@echo "  version        : Show the version of this project."
	@echo "  clean          : Remove all binary file of this project."
	@echo "  gotool         : Run go tool fmt and vet."

# 打包全部环境
# 相当于分别执行 make build 、make build-linux 、make build-windows 、make build-darwin
all: build build-linux build-windows build-darwin

# 本地环境打包：make build
build:
	@echo "\033[34mBuild the binary of this project for current platform ...\033[0m"
	go build -o ${AppName} -ldflags $(LdflagsArgs)
	@echo "\033[32mSuccess!\033[0m"

# 构建本项目可在 Linux 操作系统上执行的可执行文件
build-linux:
	@echo "\033[34mBuild the linux binary of this project ...\033[0m"
	@go mod tidy \
	&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(AppName)-linux \
	-ldflags $(LdflagsArgs)
	@echo "\033[32mSuccess!\033[0m"

# 构建本项目可在 Windows 操作系统上执行的可执行文件
build-windows:
	@echo "\033[34mBuild the windows binary of this project ...\033[0m"
	@go mod tidy \
	&& CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(AppName)-windows \
	-ldflags $(LdflagsArgs)
	@echo "\033[32mSuccess!\033[0m"

# 构建本项目可在 Mac OS 操作系统上执行的可执行文件
build-darwin:
	@echo "\033[34mBuild the darwin binary of this project ...\033[0m"
	@go mod tidy \
	&& CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(AppName)-darwin \
	-ldflags $(LdflagsArgs)
	@echo "\033[32mSuccess!\033[0m"

version:
	@echo "\033[35mThis Project Version is:\033[0m"
	@echo $(Version) \(git commit: $(GitCommit)\)

clean:
	@echo "\033[34mSearching for and deleting [$(AppName)] prefixed binary files...\033[0m"
	@# 检查文件是否可执行，这里使用 -perm +111 来代替 -executable
	@# +111 表示文件所有者、组和其他用户都有执行权限
	@#find . -type f -name "$(AppName)*" -perm +111 -print | xargs rm -f
	@find . -type f -name "$(AppName)*" -perm +111 2>/dev/null | xargs -I {} sh -c 'echo "Deleting: {}" && rm -rf {}'
	@echo "\033[32mSuccess!\033[0m"

gotool:
	@echo "\033[34mRun go tool fmt and vet ...\033[0m"
	go fmt ./...
	go vet ./...
	@echo "\033[32mSuccess!\033[0m"