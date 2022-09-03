package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao/database"
	"github.com/xilylg/accountbook/model"
)

type UserDao struct{}

var User *UserDao

func (m *UserDao) Add(c *gin.Context, User *model.User) error {
	return database.User.Add(c, User)
}

func (m *UserDao) Update(c *gin.Context, User *model.User) error {
	database.User.Update(c, User)
	return nil
}

func (m *UserDao) One(c *gin.Context, UserId int64) (*model.User, error) {
	rst, err := database.User.One(c, UserId)
	return rst, err
}

func (m *UserDao) List(c *gin.Context, uid int64, status int8, name string, page, size int) ([]*model.User, error) {
	mQuery := new(model.User)
	mQuery.Uid = uid
	mQuery.Status = status
	list, err := database.User.List(c, mQuery, page, size)
	return list, err
	return list, err
}

func (m *UserDao) Delete(c *gin.Context, UserId int64) error {
	err := database.User.Delete(c, UserId)
	return err
}
