package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao/database"
	"github.com/xilylg/accountbook/model"
)

type ItemDao struct{}

var Item *ItemDao

func (m *ItemDao) Add(c *gin.Context, Item *model.Item) error {
	return database.Item.Add(c, Item)
}

func (m *ItemDao) Update(c *gin.Context, Item *model.Item) error {
	database.Item.Update(c, Item)
	return nil
}

func (m *ItemDao) One(c *gin.Context, ItemId int64) (*model.Item, error) {
	rst, err := database.Item.One(c, ItemId)
	return rst, err
}

func (m *ItemDao) List(c *gin.Context, uid int64, status int8, name string, page, size int) ([]*model.Item, error) {
	mQuery := new(model.Item)
	mQuery.Uid = uid
	mQuery.Name = name
	mQuery.Status = status
	list, err := database.Item.List(c, mQuery, page, size)
	return list, err
}

func (m *ItemDao) Delete(c *gin.Context, itemId int64) error {
	err := database.Item.Delete(c, itemId)
	return err
}
