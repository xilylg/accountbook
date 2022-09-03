package service

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/trans"
)

func Init() {
	Member = new(MemberSvc)
	Book = new(BookSvc)
	Item = new(ItemSvc)
	User = new(UserSvc)
	Flow = new(FlowSvc)
	Target = new(TargetSvc)
}

func CheckOwner(c *gin.Context, id int64, objType string) (interface{}, trans.ECODE_T) {
	var obj interface{}
	var code trans.ECODE_T
	var uid int64
	var status int8

	switch objType {
	case "member":
		obj, code = Member.One(c, id)
		if code == trans.SUCCESS {
			uid = obj.(*model.Member).Uid
			status = obj.(*model.Member).Status
		}
	case "book":
		obj, code = Book.One(c, id)
		if code == trans.SUCCESS {
			uid = obj.(*model.Book).Uid
			status = obj.(*model.Book).Status
		}
	case "item":
		obj, code = Item.One(c, id)
		if code == trans.SUCCESS {
			uid = obj.(*model.Item).Uid
			status = obj.(*model.Item).Status
		}
	case "target":
		obj, code = Target.One(c, id)
		if code == trans.SUCCESS {
			uid = obj.(*model.Target).Uid
			status = obj.(*model.Target).Status
		}
	case "flow":
		obj, code = Flow.One(c, id)
		if code == trans.SUCCESS {
			uid = obj.(*model.Flow).Uid
			status = obj.(*model.Flow).Status
		}
	}

	if code == trans.SUCCESS {
		if status != 1 {
			return nil, trans.ERROR_RESOURCE_INVALID
		}
		if c.GetInt64("uid") != uid {
			return nil, trans.ERROR_OWNER
		}
	}
	return obj, code
}
