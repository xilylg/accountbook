package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/trans"
)

type MemberSvc struct{}

var Member *MemberSvc

func (m *MemberSvc) Add(c *gin.Context, member *model.MemberAddReq) trans.ECODE_T {
	memberDb := new(model.Member)
	memberDb.Uid = member.Uid
	memberDb.Nickname = member.Nickname
	memberDb.Relation = member.Relation
	memberDb.CreateTime = time.Now().Unix()
	memberDb.Status = 1
	err := dao.Member.Add(c, memberDb)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}

func (m *MemberSvc) Update(c *gin.Context, member *model.MemberUpdateReq) trans.ECODE_T {
	memberDb := new(model.Member)
	memberDb.MemberId = member.MemberId
	memberDb.Nickname = member.Nickname
	memberDb.Relation = member.Relation
	memberDb.UpdateTime = time.Now().Unix()
	err := dao.Member.Update(c, memberDb)
	if err == nil {
		return trans.SUCCESS
	}
	return trans.ERROR_DB_INNER
}

func (m *MemberSvc) One(c *gin.Context, mid int64) (*model.Member, trans.ECODE_T) {
	member, err := dao.Member.One(c, mid)
	if err != nil {
		return nil, trans.ERROR_INNER
	}
	return member, trans.SUCCESS
}

func (m *MemberSvc) List(c *gin.Context, mid int64, status int8, nickname string, page, size int) ([]*model.Member, trans.ECODE_T) {
	if page == 0 {
		page = 1
	}
	if size == 0 || size > 50 {
		size = 20
	}
	rst, err := dao.Member.List(c, mid, status, nickname, page, size)
	if err != nil {
		return nil, trans.ERROR_DB_INNER
	}
	return rst, trans.SUCCESS
}

func (m *MemberSvc) Delete(c *gin.Context, mid int64) trans.ECODE_T {
	err := dao.Member.Delete(c, mid)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}
