
FROM golang:1.16.3-alpine as buidler

ENV GO111MODULE=on GOPROXY=https://goproxy.cn,https://goproxy.io,direct

WORKDIR /go-project

COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk --no-cache add git curl bash tzdata \
    # set China timezone
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" >  /etc/timezone \
    && apk del tzdata \
    # && echo "machine <you-private-website-url> login <your-private-website-account> password <your-private-website-password>" > ~/.netrc  \
    && go mod tidy \
    && chmod +x ./entryPoint.sh

EXPOSE 8501

#ENTRYPOINT ["/go-project/entryPoint.sh"]