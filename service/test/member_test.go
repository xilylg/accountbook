package test

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/conf"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/log"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/validator"
)

func TestOne(t *testing.T) {
	gin.SetMode(gin.DebugMode)

	conf.Init()
	log.Init(conf.SysConf.Log)
	dao.Init(conf.SysConf)
	service.Init()

	validator.Init()

	var mid int64 = 1
	c := new(gin.Context)
	member, code := service.Member.One(c, mid)
	fmt.Printf("%+v\n%d\n", member, code)
}
