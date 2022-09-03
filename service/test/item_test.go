package test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/trans"
)

func testItemAdd(t *testing.T) {
	post := new(model.ItemAddReq)
	post.Name = "本人11"
	post.Flow = 1
	post.Uid = 3
	c := new(gin.Context)
	code := service.Item.Add(c, post)
	if code != trans.SUCCESS {
		t.Error(code)
	}
}

func testItemOne(t *testing.T) {
	c := new(gin.Context)
	rst, code := service.Item.One(c, 1)
	if code != trans.SUCCESS {
		t.Error(code)
	}
}

func TestItem(t *testing.T) {
	Init()
	t.Run("add", testItemAdd)
}
