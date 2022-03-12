package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"

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
			zap.String("request_method", c.Request.Method),                    // 当前请求的方法
			zap.String("request_path", c.Request.Host+c.Request.URL.String()), // 完整的请求地址
			zap.String("request_host", c.Request.Host),                        // 请求的 host
			zap.String("request_url", c.Request.URL.String()),                 // 请求的 url
			zap.String("request_body", string(requestBody)),                   // 请求的内容
			zap.String("client_ip", c.ClientIP()),                             // 客户端的 ip 地址
			zap.String("user-agent", c.Request.UserAgent()),                   // 用户请求头
			zap.Any("headers", c.Request.Header),                              // 请求头
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Int("response_status", responseStatus),                    // 当前的响应结果状态码
			zap.String("code_execute_time", strx.StrMicroseconds(cost)),   // 程序执行时间
			zap.String("response_body", responseBodyWriter.body.String()), // 当前的请求结果响应体
		}

		if responseStatus > 400 && responseStatus <= 499 {
			// 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
			logger.Warn("HTTP Warning [ "+cast.ToString(responseStatus)+" ]", logFields...)
		} else if responseStatus >= 500 && responseStatus <= 599 {
			// 除了内部错误，记录 error
			logger.Error("HTTP Error [ "+cast.ToString(responseStatus)+" ]", logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}

	}
}
