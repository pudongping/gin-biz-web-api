.PHONY: help all build build-linux build-windows build-darwin version clean

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

# 打包全部环境
# 相当于分别执行 make build 、make build-linux 、make build-windows 、make build-darwin
all: build build-linux build-windows build-darwin

# 本地环境打包
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
	-rm -rf $(AppName)*