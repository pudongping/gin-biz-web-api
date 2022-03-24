package middleware

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/helper/strx"
	"gin-biz-web-api/pkg/logger"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}

	return w.ResponseWriter.Write(p)
}

// AccessLog 记录请求日志
// 参考：gin.Logger()
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 如果是访问静态资源，那么不记录请求日志
		if strings.HasPrefix(c.Request.URL.Path, config.GetString("upload.static_fs_relative_path")) {
			c.Next()
			return
		}

		// 获取 response 内容
		responseBodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = responseBodyWriter

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次
			requestBody, _ = ioutil.ReadAll(c.Request.Body)
			// 读取后，重新赋值 c.Request.Body ，以供后续的其他操作
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始时间
		start := time.Now()
		c.Next()
		// 程序执行花费时间
		cost := time.Since(start)

		// http 响应状态码
		responseStatus := responseBodyWriter.Status()

		// 开始记录日志
		logFields := []zap.Field{
			zap.String("request_method", c.Request.Method),                   // 当前请求的方法
			zap.String("request_url", c.Request.Host+c.Request.URL.String()), // 完整的请求地址（host + path + query）eg：`0.0.0.0:3000/api/user?aa=11&bb=22`
			zap.String("request_path", c.Request.URL.Path),                   // 只有请求地址，不带参数 eg：`/api/user`
			zap.String("request_uri", c.Request.RequestURI),                  // 带参数的地址 eg： `/api/user?aa=11&bb=22`
			zap.String("request_query", c.Request.URL.RawQuery),              // 只有参数 eg：`aa=11&bb=22`
			// zap.String("request_body", string(requestBody)),                   // 请求的内容
			zap.String("client_ip", c.ClientIP()), // 客户端的 ip 地址
			zap.String("remote_addr", c.Request.RemoteAddr),
			zap.String("user_agent", c.Request.UserAgent()), // 用户请求头
			zap.Any("headers", c.Request.Header),            // 请求头
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Int("response_status", responseStatus),                  // 当前的响应结果状态码
			zap.String("code_execute_time", strx.StrMicroseconds(cost)), // 程序执行时间
			// zap.String("response_body", responseBodyWriter.body.String()), // 当前的请求结果响应体
		}

		// 如果是上传文件，那么则不记录请求参数内容
		if !strings.HasPrefix(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
			// 请求的内容 eg：`"x=33&y=zz"`
			logFields = append(logFields, zap.String("request_body", string(requestBody)))
		}

		// 响应的内容
		logFields = append(logFields, zap.String("response_body", responseBodyWriter.body.String()))

		// 记录访问日志
		logger.Info("HTTP Access Log [ "+cast.ToString(responseStatus)+" ]", logFields...)

	}
}
