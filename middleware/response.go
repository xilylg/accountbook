package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/utils"
)

type Response struct {
	ErrCode string      `json:"error_code"`
	ErrMsg  string      `json:"error_msg"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"trace_id"`
}

func Ok(c *gin.Context, data interface{}) {
	resp := &Response{ErrCode: "0", Data: data, TraceId: utils.GetTraceIdFromCtx(c)}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}

func Fail(c *gin.Context, code string, errmsg string) {
	traceId := ""
	if traceId, _ := c.Get("traceId"); traceId != nil {
		traceId = traceId.(string)
	}
	resp := &Response{ErrCode: "0", ErrMsg: errmsg, TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
