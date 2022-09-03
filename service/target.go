package service

import (
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/trans"
	"github.com/xilylg/accountbook/utils"
)

type TargetSvc struct{}

var Target *TargetSvc

func (m *TargetSvc) Add(c *gin.Context, target *model.TargetAddReq) trans.ECODE_T {
	TargetDb := new(model.Target)
	TargetDb.Uid = target.Uid
	TargetDb.Content = target.Content
	TargetDb.Pic = target.Pic
	TargetDb.Status = target.Status
	if target.ExpertStartTime.Unix() > 0 {
		TargetDb.ExpertStartTime = target.ExpertStartTime.Unix()
	}
	if  target.ExpertFinishTime.Unix()> 0 {
		TargetDb.ExpertFinishTime = target.ExpertFinishTime.Unix()
	}
	err := dao.Target.Add(c, TargetDb)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}

// 修改制定内容
func (m *TargetSvc) Update(c *gin.Context, target *model.TargetUpdateReq) trans.ECODE_T {
	TargetDb := new(model.Target)
	TargetDb.TargetId = target.TargetId
	TargetDb.Content = target.Content
	TargetDb.Pic = target.Pic
	if target.ExpertStartTime.Unix() > 0 {
		TargetDb.ExpertStartTime = target.ExpertStartTime.Unix()
	}
	if target.ExpertFinishTime.Unix() > 0 {
		TargetDb.ExpertFinishTime = target.ExpertFinishTime.Unix()
	}
	err := dao.Target.Update(c, TargetDb)
	if err == nil {
		return trans.SUCCESS
	}
	return trans.ERROR_DB_INNER
}

// 回顾
func (m *TargetSvc) UpdateReview(c *gin.Context, target *model.TargetReviewReq) trans.ECODE_T {
	TargetDb := new(model.Target)
	TargetDb.TargetId = target.TargetId
	TargetDb.ReviewResult = target.ReviewResult
	TargetDb.ReviewStar = target.ReviewStar
	TargetDb.ReviewTime = time.Now().Unix()
	err := dao.Target.Update(c, TargetDb)
	if err == nil {
		return trans.SUCCESS
	}
	return trans.ERROR_DB_INNER
}

func (m *TargetSvc) One(c *gin.Context, TargetId int64) (*model.Target, trans.ECODE_T) {
	Target, err := dao.Target.One(c, TargetId)
	if err != nil {
		return nil, trans.ERROR_DB_NOT_RECORD
	}

	return Target, trans.SUCCESS
}

func (m *TargetSvc) List(c *gin.Context, query *model.TargetListReq) ([]*model.Target, int64, trans.ECODE_T) {
	if query.Page == 0 {
		query.Page = 1
	}
	if query.Size == 0 || query.Size > 50 {
		query.Size = 20
	}
	target := new(model.Target)
	target.TargetId = query.TargetId
	target.Content = query.Content
	if query.ExpertStartTime.Unix() > 0 {
		target.ExpertStartTime = query.ExpertStartTime.Unix()
	}
	if query.ExpertFinishTime.Unix() > 0 {
		target.ExpertFinishTime = query.ExpertFinishTime.Unix()
	}
	target.Status = query.Status
	rst, err := dao.Target.List(c, target, query.Page, query.Size)
	total := dao.Target.Count(c, target)
	if err != nil {
		return nil, 0, trans.ERROR_DB_INNER
	}
	return rst, total, trans.SUCCESS
}

func (m *TargetSvc) Delete(c *gin.Context, TargetId int64) trans.ECODE_T {
	err := dao.Target.Delete(c, TargetId)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}
