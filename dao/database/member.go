package database

import (
	"time"

	"github.com/gin-gonic/gin"
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

func (memDb *MemberDb) Add(c *gin.Context, member *model.Member) error {
	rst := memDb.db.Create(member)
	if ProcErrorLog(c, rst, member) {
		return rst.Error
	}
	return nil
}

func (memDb *MemberDb) Update(c *gin.Context, member *model.Member) error {
	rst := memDb.db.Updates(member)
	if ProcErrorLog(c, rst, member) {
		return nil
	}
	return rst.Error
}

func (memDb *MemberDb) One(c *gin.Context, mid int64) (*model.Member, error) {
	member := new(model.Member)
	rst := memDb.db.First(member, mid)
	if ProcErrorLog(c, rst, member) {
		return nil, rst.Error
	}
	return member, nil
}

func (memDb *MemberDb) List(c *gin.Context, query *model.Member, page, size int) ([]*model.Member, error) {
	members := make([]*model.Member, 0)
	rst := memDb.db.Limit(size).Offset((page - 1) * size).Where(query).Find(&members)
	if ProcErrorLog(c, rst, map[string]interface{}{"query": query, "page": page, "size": size}) {
		return nil, rst.Error
	}
	return members, nil
}

func (memDb *MemberDb) Delete(c *gin.Context, mid int64) error {
	member := new(model.Member)
	member.MemberId = mid
	member.Status = -1
	member.UpdateTime = time.Now().Unix()
	rst := memDb.db.Updates(member)
	if ProcErrorLog(c, rst, map[string]interface{}{"mid": mid}) {
		return rst.Error
	}
	return nil
}
