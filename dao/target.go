package dao

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao/database"
	"github.com/xilylg/accountbook/model"
)

type TargetDao struct{}

var Target *TargetDao

func (m *TargetDao) Add(c *gin.Context, target *model.Target) error {
	target.CreateTime = time.Now().Unix()
	return database.Target.Add(c, target)
}

func (m *TargetDao) Update(c *gin.Context, target *model.Target) error {
	target.UpdateTime = time.Now().Unix()
	database.Target.Update(c, target)
	return nil
}

func (m *TargetDao) One(c *gin.Context, TargetId int64) (*model.Target, error) {
	rst, err := database.Target.One(c, TargetId)
	return rst, err
}

func (m *TargetDao) List(c *gin.Context, target *model.Target, page, size int) ([]*model.Target, error) {
	list, err := database.Target.List(c, target, page, size)
	return list, err
}

func (t *TargetDao) Count(c *gin.Context, target *model.Target) int64 {
	count, _ := database.Target.Count(c, target)
	return count
}

func (m *TargetDao) Delete(c *gin.Context, targetId int64) error {
	err := database.Target.Delete(c, targetId)
	return err
}
