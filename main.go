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
	//gin.ReleaseMode
	//gin.SetMode(gin.DebugMode)

	conf.Init()
	log.Init(conf.SysConf.Log)
	// defer func() {
	// 	if err := recover(); err != nil { //注意必须要判断
	// 		log.Panic(nil, "recover", err)
	// 		fmt.Printf("%s\n", err)
	// 	}
	// }() //用来调用此匿名函数

	dao.Init(conf.SysConf)
	service.Init()

	validator.Init()
	f, _ := os.Create(conf.SysConf.Base.RootPath + "/log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	rt := router.Init()

	if utils.CheckSysPort(conf.SysConf.Base.Port) {
		rt.Run(":" + conf.SysConf.Base.Port)
	} else {
		panic(conf.SysConf.Base.Port + " is invalid")
	}

}
