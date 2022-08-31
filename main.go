package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/conf"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/log"
	"github.com/xilylg/accountbook/router"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/utils"
	"github.com/xilylg/accountbook/validator"
)

func main() {
	conf.Init()
	log.Init(conf.SysConf.Log)
	dao.Init(conf.SysConf)
	validator.Init()
	service.Init()

	f, _ := os.Create(conf.SysConf.Base.RootPath + "/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	rt := router.Init()

	if utils.CheckSysPort(conf.SysConf.Base.Port) {
		rt.Run(":" + conf.SysConf.Base.Port)
	} else {
		panic(conf.SysConf.Base.Port + " is invalid")
	}

}
