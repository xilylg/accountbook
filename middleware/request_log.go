package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/log"
	"github.com/xilylg/accountbook/utils"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		RequestInLog(c)
		defer RequestOutLog(c)
		c.Next()
	}
}

func RequestInLog(c *gin.Context) {
	var traceId string
	if traceId = c.Request.Header.Get("com-header-rid"); traceId == "" {
		traceId = utils.GetTraceId()
	}
	c.Set("startExecTime", time.Now())
	c.Set("traceId", traceId)
}

func RequestOutLog(c *gin.Context) {
	endExecTime := time.Now()
	st, _ := c.Get("startExecTime")

	startExecTime, _ := st.(time.Time)
	log.Access(c, c.Request.RequestURI, map[string]interface{}{
		"method":    c.Request.Method,
		"args":      c.Request.PostForm,
		"from":      c.ClientIP(),
		"proc_time": endExecTime.Sub(startExecTime).Microseconds(),
	})
}
