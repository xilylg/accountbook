package front

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/trans"
)

type BookC struct{}

var Book *BookC

func (m *BookC) Add(c *gin.Context) {
	post := new(model.BookAddReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}
	post.Uid = c.GetInt64("uid")
	code := service.Book.Add(c, post)
	Resp(c, code, nil)
}

func (m *BookC) Update(c *gin.Context) {
	post := new(model.BookUpdateReq)
	err := c.ShouldBindWith(post, binding.Form)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	bookId, _ := strconv.ParseInt(c.Param("book_id"), 10, 64)
	_, code := service.CheckOwner(c, bookId, "book")
	fmt.Printf("%d\n", bookId)
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	post.BookId = bookId
	code = service.Book.Update(c, post)
	Resp(c, code, nil)
}

func (m *BookC) One(c *gin.Context) {
	bookId, _ := strconv.ParseInt(c.Param("book_id"), 10, 64)
	Book, code := service.CheckOwner(c, bookId, "book")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}
	Resp(c, code, Book)
}

func (m *BookC) List(c *gin.Context) {
	query := new(model.BookListReq)
	err := c.ShouldBindWith(query, binding.Query)
	if err != nil {
		Resp(c, trans.ERROR_PARAM, err.Error())
		return
	}

	rst, code := service.Book.List(c, query.BookId, 1, query.Name, query.Page, query.Size)
	Resp(c, code, rst)
}

func (m *BookC) Delete(c *gin.Context) {
	bookId, _ := strconv.ParseInt(c.Param("book_id"), 10, 64)
	_, code := service.CheckOwner(c, bookId, "book")
	if code != trans.SUCCESS {
		Resp(c, code, nil)
		return
	}

	code = service.Book.Delete(c, bookId)
	Resp(c, code, "")
}
