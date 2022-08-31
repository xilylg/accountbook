package conf

import (
	"github.com/BurntSushi/toml"

	"github.com/xilylg/accountbook/dao/database"
	"github.com/xilylg/accountbook/log"
	"github.com/xilylg/accountbook/utils"
)

type Base struct {
	Port            string
	RootPath        string
	DefaultLanguage string `toml:"default_language"`
	TimeLocation    string `toml:"time_location"`
}

type SystemConf struct {
	Base     Base
	Log      *log.LogConf
	Database *database.DatabaseConf
}

const (
	EVN string = "dev"
)

var SysConf *SystemConf

func Init() {
	rootPath := utils.CurrentRootPath()
	SysConf = new(SystemConf)
	path := rootPath + "/conf/" + EVN + "/server.toml"
	if _, err := toml.DecodeFile(path, SysConf); err != nil {
		panic("sysconf setup fail " + (err.Error()))
	}

	// 更正文件夹路径
	SysConf.Base.RootPath = rootPath
	SysConf.Log.Path = rootPath + "/" + SysConf.Log.Path
	return
}
