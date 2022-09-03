package database

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/model"
	"gorm.io/gorm"
)

type BookDb struct {
	db *gorm.DB
}

var Book *BookDb

func registerBook(db *gorm.DB) *BookDb {
	Book := new(BookDb)
	Book.db = db
	return Book
}

func (memDb *BookDb) Add(c *gin.Context, Book *model.Book) error {
	rst := memDb.db.Create(Book)
	if ProcErrorLog(c, rst, Book) {
		return rst.Error
	}
	return nil
}

func (memDb *BookDb) Update(c *gin.Context, Book *model.Book) error {
	rst := memDb.db.Updates(Book)
	if ProcErrorLog(c, rst, Book) {
		return nil
	}
	return rst.Error
}

func (memDb *BookDb) One(c *gin.Context, bookId int64) (*model.Book, error) {
	Book := new(model.Book)
	rst := memDb.db.First(Book, bookId)
	if ProcErrorLog(c, rst, Book) {
		return nil, rst.Error
	}
	return Book, nil
}

func (memDb *BookDb) List(c *gin.Context, query *model.Book, page, size int) ([]*model.Book, error) {
	Books := make([]*model.Book, 0)
	rst := memDb.db.Limit(size).Offset((page - 1) * size).Where(query).Find(&Books)
	if ProcErrorLog(c, rst, map[string]interface{}{"query": query, "page": page, "size": size}) {
		return nil, rst.Error
	}
	return Books, nil
}

func (memDb *BookDb) Delete(c *gin.Context, bookId int64) error {
	Book := new(model.Book)
	Book.BookId = bookId
	Book.Status = -1
	Book.UpdateTime = time.Now().Unix()
	rst := memDb.db.Updates(Book)
	if ProcErrorLog(c, rst, map[string]interface{}{"bookId": bookId}) {
		return rst.Error
	}
	return nil
}
