package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/log"
	"github.com/xilylg/accountbook/model"
)

type MemberSvc struct{}

var Member *MemberSvc

func (m *MemberSvc) Add(c *gin.Context, member *model.MemberAddReq) error {
	memberDb := new(model.Member)
	memberDb.Uid = member.Uid
	memberDb.Nickname = member.Nickname
	memberDb.Relation = member.Relation
	err := dao.Member.Add(memberDb)
	if err != nil {
		log.Error(c, "add_member_fail", map[string]interface{}{"error": err.Error()})
	}
	return err
}
