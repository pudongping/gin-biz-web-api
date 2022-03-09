package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
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
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 response 内容
		ResponseBodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = ResponseBodyWriter

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
		responseStatus := c.Writer.Status()

		// 开始记录日志
		logFields := []zap.Field{
			zap.String("request_method", c.Request.Method),
			zap.String("request_url", c.Request.URL.String()),
			zap.String("request_query", c.Request.URL.RawQuery),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Int("response_status", responseStatus),
			zap.String("code_execute_time", strx.StrMicroseconds(cost)),
		}

		logger.Debug("HTTP Access Log", logFields...)

	}
}
