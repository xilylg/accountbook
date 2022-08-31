package front

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xilylg/accountbook/middleware"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/service"
)

type MemberC struct{}

var Member *MemberC

func (m *MemberC) Add(c *gin.Context) {
	var post *model.MemberAddReq
	if err := c.ShouldBindWith(post, binding.Query); err == nil {
		service.Member.Add(c, post)
		middleware.Ok(c, "admin pong")
	} else {
		middleware.Fail(c, "100", err.Error())
	}
}
