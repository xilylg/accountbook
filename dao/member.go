package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao/database"
	"github.com/xilylg/accountbook/model"
)

type MemberDao struct{}

var Member *MemberDao

func (m *MemberDao) Add(c *gin.Context, member *model.Member) error {
	return database.Member.Add(c, member)
}

func (m *MemberDao) Update(c *gin.Context, member *model.Member) error {
	database.Member.Update(c, member)
	return nil
}

func (m *MemberDao) One(c *gin.Context, mid int64) (*model.Member, error) {
	rst, err := database.Member.One(c, mid)
	return rst, err
}

func (m *MemberDao) List(c *gin.Context, uid int64, status int8, nickname string, page, size int) ([]*model.Member, error) {
	mQuery := new(model.Member)
	mQuery.Uid = uid
	mQuery.Nickname = nickname
	mQuery.Status = status
	list, err := database.Member.List(c, mQuery, page, size)
	return list, err
}

func (m *MemberDao) Delete(c *gin.Context, mid int64) error {
	err := database.Member.Delete(c, mid)
	return err
}
