package front

import (
	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/service"
	"github.com/xilylg/accountbook/trans"
)

func CheckOwner(c *gin.Context, id int64, objType string) (interface{}, trans.ECODE_T) {
	var obj interface{}
	var code trans.ECODE_T
	var uid int64
	var status int8
	switch objType {
	case "member":
		obj, code = service.Member.One(c, id)
		uid = obj.(*model.Member).Uid
		status = obj.(*model.Member).Status
	}

	if code == trans.SUCCESS {
		if c.GetInt64("uid") != uid {
			return nil, trans.ERROR_OWNER
		}
		if status != 1 {
			return nil, trans.ERROR_RESOURCE_INVALID
		}
	}
	return obj, code
}
