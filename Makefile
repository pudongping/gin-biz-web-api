.PHONY: build-local build-server version help clean

AppName='gin-biz-web-api'
# v1.2.0-11-g98d9b66-dev
Version := $(shell git describe --tags --always --dirty="-dev")
# 98d9b668dd382b6f0f110782667a3ebf882109ee
GitCommit := $(shell git rev-list -1 HEAD)

# 本地打包
build-local:
	@echo 'build local server binary file'
	go build -ldflags \
	"-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X 'main.goVersion=`go version`' -X main.gitCommitID=`git rev-parse HEAD`"

# Linux 服务器上打包
build-server:
	@echo 'build linux server binary file'
	go mod tidy \
	&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(AppName) \
	-ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X 'main.goVersion=`go version`' -X main.gitCommitID=`git rev-parse HEAD`"

version:
	@echo 'This Project Version is:'
	@echo $(Version) \(git commit: $(GitCommit)\)

help:
	@echo "use \033[1;31;44m [ make ] \033[0m or [ make build-local ] command: build local server binary file."
	@echo "use [ make build-server ] command: build linux server binary file."
	@echo "use [ make version ] command: print this project version."
	@echo "use [ make clean ] command: remove this object binary file."

clean:
	-rm -rf $(AppName)