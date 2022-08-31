package test

import (
	"testing"

	"github.com/xilylg/accountbook/conf"
	"github.com/xilylg/accountbook/dao/database"
	"github.com/xilylg/accountbook/model"
)

func TestExam(t *testing.T) {
	sysConf := conf.Init()
	db, err := database.Init(sysConf.Database)
	if err != nil {
		t.Error(err)
	}
	us := &model.User{Nickname: "aaa", Password: "aaaa"}
	err = db.User.Add(us)
	if err != nil {
		t.Error(err)
	}
}
