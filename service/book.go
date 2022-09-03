package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/trans"
)

type BookSvc struct{}

var Book *BookSvc

func (m *BookSvc) Add(c *gin.Context, Book *model.BookAddReq) trans.ECODE_T {
	BookDb := new(model.Book)
	BookDb.Uid = Book.Uid
	BookDb.Name = Book.Name
	BookDb.Cover = Book.Cover
	BookDb.TargetId = Book.TargetId
	BookDb.CycleType = Book.CycleType
	BookDb.CycleStart = Book.CycleStart
	BookDb.CycleEnd = Book.CycleEnd
	BookDb.CreateTime = time.Now().Unix()
	BookDb.Status = 1
	err := dao.Book.Add(c, BookDb)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}

func (m *BookSvc) Update(c *gin.Context, Book *model.BookUpdateReq) trans.ECODE_T {
	BookDb := new(model.Book)
	BookDb.BookId = Book.BookId
	BookDb.Name = Book.Name
	BookDb.Cover = Book.Cover
	BookDb.TargetId = Book.TargetId
	BookDb.CycleType = Book.CycleType
	BookDb.CycleStart = Book.CycleStart
	BookDb.CycleEnd = Book.CycleEnd
	BookDb.UpdateTime = time.Now().Unix()
	err := dao.Book.Update(c, BookDb)
	if err == nil {
		return trans.SUCCESS
	}
	return trans.ERROR_DB_INNER
}

func (m *BookSvc) One(c *gin.Context, bookId int64) (*model.Book, trans.ECODE_T) {
	book, err := dao.Book.One(c, bookId)
	if err != nil {
		return nil, trans.ERROR_DB_NOT_RECORD
	}

	return book, trans.SUCCESS
}

func (m *BookSvc) List(c *gin.Context, bookId int64, status int8, name string, page, size int) ([]*model.Book, trans.ECODE_T) {
	if page == 0 {
		page = 1
	}
	if size == 0 || size > 50 {
		size = 20
	}
	rst, err := dao.Book.List(c, bookId, status, name, page, size)
	if err != nil {
		return nil, trans.ERROR_DB_INNER
	}
	return rst, trans.SUCCESS
}

func (m *BookSvc) Delete(c *gin.Context, bookId int64) trans.ECODE_T {
	err := dao.Book.Delete(c, bookId)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}
