package front

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/trans"
)

type UserC struct{}

var User *UserC

func (m *UserC) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (m *UserC) Update(c *gin.Context) {
	post := new(model.UserUpdateReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}
	post.Uid = c.GetInt64("uid")
	code := service.User.Update(c, post)
	Resp(c, code, nil)
}

func (m *UserC) One(c *gin.Context) {
	uid := c.GetInt64("uid")
	user, code := service.User.One(c, uid)
	Resp(c, code, user)
}
