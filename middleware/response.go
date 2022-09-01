package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/trans"
	"github.com/xilylg/accountbook/utils"
)

type Response struct {
	ErrCode trans.ECODE_T `json:"error_code"`
	ErrMsg  string        `json:"error_msg"`
	Data    interface{}   `json:"data"`
	TraceId string        `json:"trace_id"`
}

func Ok(c *gin.Context, data interface{}) {
	resp := &Response{ErrCode: trans.SUCCESS, Data: data, TraceId: utils.GetTraceIdFromCtx(c)}
	c.JSON(200, resp)
}

func Fail(c *gin.Context, code trans.ECODE_T, data interface{}) {
	traceId := ""
	if traceId, _ := c.Get("traceId"); traceId != nil {
		traceId = traceId.(string)
	}
	resp := &Response{ErrCode: code, ErrMsg: trans.ERRORMSG[code], TraceId: traceId, Data: data}
	c.JSON(200, resp)
}

func Resp(c *gin.Context, code trans.ECODE_T, data interface{}) {
	if code == trans.SUCCESS {
		Ok(c, data)
	} else {
		Fail(c, code, data)
	}
}
