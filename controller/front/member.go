package front

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xilylg/accountbook/middleware"
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
		middleware.Fail(c, trans.ERROR_INNER, err.Error())
		return
	}
	code := service.Member.Add(c, post)
	middleware.Resp(c, code, nil)
}

func (m *MemberC) Update(c *gin.Context) {
	post := new(model.MemberUpdateReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		middleware.Fail(c, trans.ERROR_PARAM, err.Error())
	}

	mid, _ := strconv.ParseInt(c.Param("mid"), 10, 64)
	_, code := CheckOwner(c, mid, "member")
	if code != trans.SUCCESS {
		middleware.Resp(c, code, nil)
		return
	}
	post.MemberId = mid
	code = service.Member.Update(c, post)
	middleware.Resp(c, code, nil)
}

func (m *MemberC) One(c *gin.Context) {
	mid, _ := strconv.ParseInt(c.Param("mid"), 10, 64)
	member, code := CheckOwner(c, mid, "member")
	if code != trans.SUCCESS {
		middleware.Resp(c, code, nil)
		return
	}
	middleware.Resp(c, code, member)
}

func (m *MemberC) List(c *gin.Context) {
	query := new(model.MemberListReq)
	err := c.ShouldBindWith(query, binding.Query)
	if err != nil {
		middleware.Fail(c, trans.ERROR_PARAM, err.Error())
		return
	}

	rst, code := service.Member.List(c, query.MemberId, 1, query.Nickname, query.Page, query.Size)
	middleware.Resp(c, code, rst)
}

func (m *MemberC) Delete(c *gin.Context) {
	mid, _ := strconv.ParseInt(c.Param("mid"), 10, 64)
	_, code := CheckOwner(c, mid, "member")
	if code != trans.SUCCESS {
		middleware.Resp(c, code, nil)
		return
	}

	code = service.Member.Delete(c, mid)
	middleware.Resp(c, code, "")
}
