package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao/database"
	"github.com/xilylg/accountbook/model"
)

type FlowDao struct{}

var Flow *FlowDao

func (m *FlowDao) Add(c *gin.Context, Flow *model.Flow) error {
	return database.Flow.Add(c, Flow)
}

func (m *FlowDao) Update(c *gin.Context, Flow *model.Flow) error {
	database.Flow.Update(c, Flow)
	return nil
}

func (m *FlowDao) One(c *gin.Context, FlowId int64) (*model.Flow, error) {
	rst, err := database.Flow.One(c, FlowId)
	return rst, err
}

func (m *FlowDao) List(c *gin.Context, flow *model.Flow, page, size int) ([]*model.Flow, error) {
	list, err := database.Flow.List(c, flow, page, size)
	return list, err
}

func (f *FlowDao) Count(c *gin.Context, flow *model.Flow) ( int64,  error) {
	return database.Flow.Count(c, flow)
}

func (m *FlowDao) Delete(c *gin.Context, FlowId int64) error {
	err := database.Flow.Delete(c, FlowId)
	return err
}
