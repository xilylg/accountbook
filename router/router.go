package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/middleware"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.AccessLog())
	adminRouter(r)
	frontRouter(r)
	return r
}
