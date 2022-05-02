#!/bin/sh

cd /go-project

app_name=gin-biz-web-api
log_path=storage/logs

if [ ! -d ${log_path} ];then
  mkdir -p ${log_path}
fi

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${app_name}

./${app_name} >> "${log_path}/std-`date +"%Y-%m-%d"`.log" 2>&1