package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/dao"
	"github.com/xilylg/accountbook/model"
	"github.com/xilylg/accountbook/trans"
)

type UserSvc struct{}

var User *UserSvc

func (m *UserSvc) Add(c *gin.Context, user *model.UserAddReq) trans.ECODE_T {
	UserDb := new(model.User)
	UserDb.Nickname = user.Nickname
	UserDb.Password = user.Password
	UserDb.Gender = user.Gender
	UserDb.Birthday = user.Birthday
	UserDb.Status = 1
	UserDb.CreateTime = time.Now().Unix()
	err := dao.User.Add(c, UserDb)
	if err != nil {
		return trans.ERROR_DB_INNER
	}
	return trans.SUCCESS
}

// 修改制定内容
func (m *UserSvc) Update(c *gin.Context, user *model.UserUpdateReq) trans.ECODE_T {
	UserDb := new(model.User)
	UserDb.Uid = user.Uid
	UserDb.Nickname = user.Nickname
	UserDb.Password = user.Password
	UserDb.Gender = user.Gender
	UserDb.Birthday = user.Birthday
	UserDb.UpdateTime = time.Now().Unix()
	err := dao.User.Update(c, UserDb)
	if err == nil {
		return trans.SUCCESS
	}
	return trans.ERROR_DB_INNER
}

func (m *UserSvc) One(c *gin.Context, userId int64) (*model.User, trans.ECODE_T) {
	User, err := dao.User.One(c, userId)
	if err != nil {
		return nil, trans.ERROR_DB_NOT_RECORD
	}

	return User, trans.SUCCESS
}
