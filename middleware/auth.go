package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/trans"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var uid int64 = 3
		if uid == 0 {
			c.AbortWithError(int(trans.ERROR_NOLOGIN), errors.New(trans.ERRORMSG[trans.ERROR_NOLOGIN]))
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
