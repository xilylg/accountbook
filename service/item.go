package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/trans"
)

type ItemSvc struct{}

var Item *ItemSvc

func (m *ItemSvc) Add(c *gin.Context, Item *model.ItemAddReq) trans.ECODE_T {
	ItemDb := new(model.Item)
	ItemDb.Uid = Item.Uid
	ItemDb.Name = Item.Name
	ItemDb.Flow = Item.Flow
	ItemDb.Status = 1
	ItemDb.CreateTime = time.Now().Unix()
	err := dao.Item.Add(c, ItemDb)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}

func (m *ItemSvc) Update(c *gin.Context, Item *model.ItemUpdateReq) trans.ECODE_T {
	ItemDb := new(model.Item)
	ItemDb.ItemId = Item.ItemId
	ItemDb.Uid = Item.Uid
	ItemDb.Name = Item.Name
	ItemDb.Flow = Item.Flow
	ItemDb.UpdateTime = time.Now().Unix()
	err := dao.Item.Update(c, ItemDb)
	if err == nil {
		return trans.SUCCESS
	}
	return trans.ERROR_DB_INNER
}

func (m *ItemSvc) One(c *gin.Context, ItemId int64) (*model.Item, trans.ECODE_T) {
	Item, err := dao.Item.One(c, ItemId)
	if err != nil {
		return nil, trans.ERROR_DB_NOT_RECORD
	}

	return Item, trans.SUCCESS
}

func (m *ItemSvc) List(c *gin.Context, ItemId int64, status int8, name string, page, size int) ([]*model.Item, trans.ECODE_T) {
	if page == 0 {
		page = 1
	}
	if size == 0 || size > 50 {
		size = 20
	}
	rst, err := dao.Item.List(c, ItemId, status, name, page, size)
	if err != nil {
		return nil, trans.ERROR_DB_INNER
	}
	return rst, trans.SUCCESS
}

func (m *ItemSvc) Delete(c *gin.Context, itemId int64) trans.ECODE_T {
	err := dao.Item.Delete(c, itemId)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}
