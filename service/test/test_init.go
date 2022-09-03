package test

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/conf"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/log"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/validator"
)

func Init() {
	gin.SetMode(gin.DebugMode)

	conf.Init()
	log.Init(conf.SysConf.Log)
	dao.Init(conf.SysConf)
	service.Init()
	validator.Init()
}
