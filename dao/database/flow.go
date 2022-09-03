package database

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/model"
	"gorm.io/gorm"
)

type FlowDb struct {
	db *gorm.DB
}

var Flow *FlowDb

func registerFlow(db *gorm.DB) *FlowDb {
	Flow := new(FlowDb)
	Flow.db = db
	return Flow
}

// func (db *FlowDb.db) TableName() string {
// 	return "flow_" + time.Now().YEAR()
// }

func (FlowDb *FlowDb) Add(c *gin.Context, flow *model.Flow) error {
	rst := FlowDb.db.Create(flow)
	if ProcErrorLog(c, rst, flow) {
		return rst.Error
	}
	return nil
}

func (FlowDb *FlowDb) Update(c *gin.Context, flow *model.Flow) error {
	rst := FlowDb.db.Updates(flow)
	if ProcErrorLog(c, rst, flow) {
		return nil
	}
	return rst.Error
}

func (FlowDb *FlowDb) One(c *gin.Context, flowId int64) (*model.Flow, error) {
	flow := new(model.Flow)
	rst := FlowDb.db.First(flow, flowId)
	if ProcErrorLog(c, rst, flow) {
		return nil, rst.Error
	}
	return flow, nil
}

func (FlowDb *FlowDb) List(c *gin.Context, query *model.Flow, page, size int) ([]*model.Flow, error) {
	Flows := make([]*model.Flow, 0)
	rst := FlowDb.db.Limit(size).Offset((page - 1) * size).Where(query).Find(&Flows)
	if ProcErrorLog(c, rst, map[string]interface{}{"query": query, "page": page, "size": size}) {
		return nil, rst.Error
	}
	return Flows, nil
}

func (FlowDb *FlowDb) Count(c *gin.Context, query *model.Flow) (total int64, err error) {
	rst := FlowDb.db.Where(query).Count(&total)
	if ProcErrorLog(c, rst, query) {
		err = rst.Error
	}
	return
}

func (FlowDb *FlowDb) Delete(c *gin.Context, flowId int64) error {
	flow := new(model.Flow)
	flow.FlowId = flowId
	flow.Status = -1
	flow.UpdateTime = time.Now().Unix()
	rst := FlowDb.db.Updates(flow)
	if ProcErrorLog(c, rst, map[string]interface{}{"flowId": flowId}) {
		return rst.Error
	}
	return nil
}
