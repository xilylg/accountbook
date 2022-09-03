package front

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/trans"
)

type ItemC struct{}

var Item *ItemC

func (m *ItemC) Add(c *gin.Context) {
	post := new(model.ItemAddReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}
	post.Uid = c.GetInt64("uid")
	code := service.Item.Add(c, post)
	Resp(c, code, nil)
}

func (m *ItemC) Update(c *gin.Context) {
	post := new(model.ItemUpdateReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	itemId, _ := strconv.ParseInt(c.Param("item_id"), 10, 64)
	_, code := service.CheckOwner(c, itemId, "item")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	post.ItemId = itemId
	code = service.Item.Update(c, post)
	Resp(c, code, nil)
}

func (m *ItemC) One(c *gin.Context) {
	ItemId, _ := strconv.ParseInt(c.Param("item_id"), 10, 64)
	Item, code := service.CheckOwner(c, ItemId, "item")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	Resp(c, code, Item)
}

func (m *ItemC) List(c *gin.Context) {
	query := new(model.ItemListReq)
	err := c.ShouldBindWith(query, binding.Query)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	rst, code := service.Item.List(c, query.ItemId, 1, query.Name, query.Page, query.Size)
	Resp(c, code, rst)
}

func (m *ItemC) Delete(c *gin.Context) {
	itemId, _ := strconv.ParseInt(c.Param("item_id"), 10, 64)
	_, code := service.CheckOwner(c, itemId, "item")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}

	code = service.Item.Delete(c, itemId)
	Resp(c, code, "")
}
