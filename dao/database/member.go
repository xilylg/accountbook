package database

import (
	"github.com/xilylg/accountbook/model"
	"gorm.io/gorm"
)

type MemberDb struct {
	db *gorm.DB
}

var Member *MemberDb

func registerMember(db *gorm.DB) *MemberDb {
	member := new(MemberDb)
	member.db = db
	return member
}

func (memDb *MemberDb) Add(member *model.Member) error {
	result := memDb.db.Create(member)
	return result.Error
}
