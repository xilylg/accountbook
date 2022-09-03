package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/trans"
)

type FlowSvc struct{}

var Flow *FlowSvc

func (m *FlowSvc) Add(c *gin.Context, Flow *model.FlowAddReq) trans.ECODE_T {
	FlowDb := new(model.Flow)
	FlowDb.Uid = Flow.Uid
	FlowDb.Cost = Flow.Cost
	FlowDb.BookId = Flow.BookId
	FlowDb.MemberId = Flow.MemberId
	FlowDb.Tip = Flow.Tip
	FlowDb.Status = 1
	FlowDb.CreateTime = time.Now().Unix()
	err := dao.Flow.Add(c, FlowDb)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}

// 修改制定内容
func (m *FlowSvc) Update(c *gin.Context, Flow *model.FlowUpdateReq) trans.ECODE_T {
	FlowDb := new(model.Flow)
	FlowDb.FlowId = Flow.FlowId
	FlowDb.Cost = Flow.Cost
	FlowDb.BookId = Flow.BookId
	FlowDb.MemberId = Flow.MemberId
	FlowDb.Tip = Flow.Tip
	FlowDb.UpdateTime = time.Now().Unix()
	err := dao.Flow.Update(c, FlowDb)
	if err == nil {
		return trans.SUCCESS
	}
	return trans.ERROR_DB_INNER
}

// 回顾
func (m *FlowSvc) UpdateReview(c *gin.Context, flowQuery *model.FlowReviewReq) trans.ECODE_T {
	flow := new(model.Flow)
	flow.FlowId = flowQuery.FlowId
	flow.ReviewTime = time.Now().Unix()
	flow.Tip = flowQuery.Tip
	flow.ReviewStar = flowQuery.ReviewStar
	flow.UpdateTime = time.Now().Unix()
	err := dao.Flow.Update(c, flow)
	if err == nil {
		return trans.SUCCESS
	}
	return trans.ERROR_DB_INNER
}

func (m *FlowSvc) One(c *gin.Context, FlowId int64) (*model.Flow, trans.ECODE_T) {
	Flow, err := dao.Flow.One(c, FlowId)
	if err != nil {
		return nil, trans.ERROR_DB_NOT_RECORD
	}

	return Flow, trans.SUCCESS
}

func (m *FlowSvc) List(c *gin.Context, query *model.FlowListReq) ([]*model.Flow, int64, trans.ECODE_T) {
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Size == 0 || query.Size > 50 {
		query.Size = 20
	}
	Flow := new(model.Flow)
	Flow.FlowId = query.FlowId
	Flow.BookId = query.BookId
	Flow.MemberId = query.MemberId
	Flow.Status = query.Status
	rst, err1 := dao.Flow.List(c, Flow, query.Page, query.Size)
	total, err2 := dao.Flow.Count(c, Flow)
	if err1 != nil || err2 != nil {
		return nil, 0, trans.ERROR_DB_INNER
	}
	return rst, total, trans.SUCCESS
}

func (m *FlowSvc) Delete(c *gin.Context, FlowId int64) trans.ECODE_T {
	err := dao.Flow.Delete(c, FlowId)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}
