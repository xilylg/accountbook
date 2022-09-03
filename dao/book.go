package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao/database"
	"github.com/xilylg/accountbook/model"
)

type BookDao struct{}

var Book *BookDao

func (m *BookDao) Add(c *gin.Context, Book *model.Book) error {
	return database.Book.Add(c, Book)
}

func (m *BookDao) Update(c *gin.Context, Book *model.Book) error {
	database.Book.Update(c, Book)
	return nil
}

func (m *BookDao) One(c *gin.Context, bookId int64) (*model.Book, error) {
	rst, err := database.Book.One(c, bookId)
	return rst, err
}

func (m *BookDao) List(c *gin.Context, uid int64, status int8, name string, page, size int) ([]*model.Book, error) {
	mQuery := new(model.Book)
	mQuery.Uid = uid
	mQuery.Name = name
	mQuery.Status = status
	list, err := database.Book.List(c, mQuery, page, size)
	return list, err
}

func (m *BookDao) Delete(c *gin.Context, bookId int64) error {
	err := database.Book.Delete(c, bookId)
	return err
}
