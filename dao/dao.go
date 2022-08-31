package dao

import (
	"github.com/xilylg/accountbook/conf"
	"github.com/xilylg/accountbook/dao/database"
)

func Init(sysConf *conf.SystemConf) {
	_, err := database.Init(sysConf.Database)
	if err != nil {
		panic("database init fail")
	}
	// todo 初始化缓存

	//初始化各服务
	Member = new(MemberDao)
}
