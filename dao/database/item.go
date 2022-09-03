package database

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/model"
	"gorm.io/gorm"
)

type ItemDb struct {
	db *gorm.DB
}

var Item *ItemDb

func registerItem(db *gorm.DB) *ItemDb {
	Item := new(ItemDb)
	Item.db = db
	return Item
}

func (itemDb *ItemDb) Add(c *gin.Context, item *model.Item) error {
	rst := itemDb.db.Create(item)
	if ProcErrorLog(c, rst, item) {
		return rst.Error
	}
	return nil
}

func (itemDb *ItemDb) Update(c *gin.Context, item *model.Item) error {
	rst := itemDb.db.Updates(item)
	if ProcErrorLog(c, rst, item) {
		return nil
	}
	return rst.Error
}

func (itemDb *ItemDb) One(c *gin.Context, itemId int64) (*model.Item, error) {
	item := new(model.Item)
	rst := itemDb.db.First(item, itemId)
	if ProcErrorLog(c, rst, item) {
		return nil, rst.Error
	}
	return item, nil
}

func (itemDb *ItemDb) List(c *gin.Context, query *model.Item, page, size int) ([]*model.Item, error) {
	items := make([]*model.Item, 0)
	rst := itemDb.db.Limit(size).Offset((page - 1) * size).Where(query).Find(&items)
	if ProcErrorLog(c, rst, map[string]interface{}{"query": query, "page": page, "size": size}) {
		return nil, rst.Error
	}
	return items, nil
}

func (itemDb *ItemDb) Delete(c *gin.Context, itemId int64) error {
	item := new(model.Item)
	item.ItemId = itemId
	item.Status = -1
	item.UpdateTime = time.Now().Unix()
	rst := itemDb.db.Updates(item)
	if ProcErrorLog(c, rst, map[string]interface{}{"itemId": itemId}) {
		return rst.Error
	}
	return nil
}
