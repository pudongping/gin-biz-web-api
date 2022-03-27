# 本地打包
.PHONY: build-local
build-local:
	go build -ldflags \
	"-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X 'main.goVersion=`go version`' -X main.gitCommitID=`git rev-parse HEAD`"