package database

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/model"
	"gorm.io/gorm"
)

type TargetDb struct {
	db *gorm.DB
}

var Target *TargetDb

func registerTarget(db *gorm.DB) *TargetDb {
	Target := new(TargetDb)
	Target.db = db
	return Target
}

func (TargetDb *TargetDb) Add(c *gin.Context, Target *model.Target) error {
	rst := TargetDb.db.Create(Target)
	if ProcErrorLog(c, rst, Target) {
		return rst.Error
	}
	return nil
}

func (TargetDb *TargetDb) Update(c *gin.Context, Target *model.Target) error {
	rst := TargetDb.db.Updates(Target)
	if ProcErrorLog(c, rst, Target) {
		return nil
	}
	return rst.Error
}

func (TargetDb *TargetDb) One(c *gin.Context, TargetId int64) (*model.Target, error) {
	Target := new(model.Target)
	rst := TargetDb.db.First(Target, TargetId)
	if ProcErrorLog(c, rst, Target) {
		return nil, rst.Error
	}
	return Target, nil
}

func (TargetDb *TargetDb) List(c *gin.Context, query *model.Target, page, size int) ([]*model.Target, error) {
	Targets := make([]*model.Target, 0)
	rst := TargetDb.db.Limit(size).Offset((page - 1) * size).Where(query).Find(&Targets)
	if ProcErrorLog(c, rst, map[string]interface{}{"query": query, "page": page, "size": size}) {
		return nil, rst.Error
	}
	return Targets, nil
}

func (TargetDb *TargetDb) Count(c *gin.Context, query *model.Target) (total int64, err error) {
	rst := TargetDb.db.Where(query).Count(&total)
	if ProcErrorLog(c, rst, query) {
		err = rst.Error
	}
	return
}

func (TargetDb *TargetDb) Delete(c *gin.Context, targetId int64) error {
	Target := new(model.Target)
	Target.TargetId = targetId
	Target.Status = -1
	Target.UpdateTime = time.Now().Unix()
	rst := TargetDb.db.Updates(Target)
	if ProcErrorLog(c, rst, map[string]interface{}{"targetId": targetId}) {
		return rst.Error
	}
	return nil
}
