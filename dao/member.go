package dao

import (
	"github.com/xilylg/accountbook/dao/database"
	"github.com/xilylg/accountbook/model"
)

type MemberDao struct{}

var Member *MemberDao

func (m *MemberDao) Add(member *model.Member) error {
	return database.Member.Add(member)
}
