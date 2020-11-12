package middleware

import (
	"blog/global"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

/**
访问日志记录中间件
记录请求的
方法、参数、响应结果、状态码、请求的开始时间和结束时间
*/
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

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()
		global.Logger.WithFields(logrus.Fields{
			//"request":  c.Request.PostForm.Encode(),
			//"response": bodyWriter.body.String(), //这是请求得返回数据，通过重写Writer接口取得
			"url":       c.Request.URL.Path,
			"method":    c.Request.Method,
			"httpCode":  bodyWriter.Status(),
			"beginTime": beginTime,
			"endTime":   endTime,
			"questTime": endTime - beginTime,
			"ip":        c.Request.Host,
		}).Info("Request_Log")
	}
}
