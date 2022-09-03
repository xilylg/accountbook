package front

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/trans"
)

type MemberC struct{}

var Member *MemberC

func (m *MemberC) Add(c *gin.Context) {
	post := new(model.MemberAddReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}
	post.Uid = c.GetInt64("uid")
	code := service.Member.Add(c, post)
	Resp(c, code, nil)
}

func (m *MemberC) Update(c *gin.Context) {
	post := new(model.MemberUpdateReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	mid, _ := strconv.ParseInt(c.Param("mid"), 10, 64)
	_, code := service.CheckOwner(c, mid, "member")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	post.MemberId = mid
	code = service.Member.Update(c, post)
	Resp(c, code, nil)
}

func (m *MemberC) One(c *gin.Context) {
	mid, _ := strconv.ParseInt(c.Param("mid"), 10, 64)
	member, code :=  service.CheckOwner(c, mid, "member")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	Resp(c, code, member)
}

func (m *MemberC) List(c *gin.Context) {
	query := new(model.MemberListReq)
	err := c.ShouldBindWith(query, binding.Query)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	rst, code := service.Member.List(c, query.MemberId, 1, query.Nickname, query.Page, query.Size)
	Resp(c, code, rst)
}

func (m *MemberC) Delete(c *gin.Context) {
	mid, _ := strconv.ParseInt(c.Param("mid"), 10, 64)
	_, code :=  service.CheckOwner(c, mid, "member")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}

	code = service.Member.Delete(c, mid)
	Resp(c, code, "")
}
