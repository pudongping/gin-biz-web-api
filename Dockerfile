
FROM golang:1.16.3-alpine as builder

LABEL maintainer="Alex <276558492@qq.com>"

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 设置工作目录为 /go-project，之后的指令都在这个目录下执行
WORKDIR /go-project

# 将当前上下文中的文件复制到容器的 /go-project 目录
COPY . .

RUN go mod tidy \
    # 将我们的代码编译成二进制可执行文件 gin-biz-web-api
    && go build -a -installsuffix cgo -o gin-biz-web-api .


###################
# 接下来创建一个小镜像
###################
FROM alpine:3.20

# 变量 timezone，可以在构建镜像时通过 --build-arg 参数传入值，例如：
# docker build --build-arg timezone=Asia/Shanghai -t gin-biz-web-api:v1.0.0 .
ARG timezone

# 设置环境变量 TIMEZONE，默认值为 Asia/Shanghai，或者使用传入的 timezone 参数值
ENV TIMEZONE=${timezone:-"Asia/Shanghai"} \
    GIN_MODE=debug

WORKDIR /go-project-run

COPY ./wait-for.sh ./
COPY ./etc/config.yaml.example ./etc/config.yaml

# 从 builder 镜像中把编译好的可执行文件拷贝到当前目录
COPY --from=builder /go-project/gin-biz-web-api ./

RUN set -eux \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk --no-cache add curl bash tzdata \
    # - config timezone
    && ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
#    如果需要删掉 tzdata 的话，就需要使用 cp 而不是通过软连接 ln 的方式，这里不删的原因是可能会用到其它时区
#    && cp /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
    && echo "${TIMEZONE}" > /etc/timezone \
#    && apk del tzdata \
    # ---------- clear works ----------
    && rm -rf /var/cache/apk/* /tmp/* /usr/share/man \
    && chmod 755 wait-for.sh \
    && echo -e "\033[42;37m Build Completed :).\033[0m\n"
