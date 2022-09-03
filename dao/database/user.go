package database

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/model"
	"gorm.io/gorm"
)

// // 自定义表名
// func (UserDb) TableName() string {
// 	return "user"
// }
type UserDb struct {
	db *gorm.DB
}

var User *UserDb

func registerUser(db *gorm.DB) *UserDb {
	User := new(UserDb)
	User.db = db
	return User
}

func (UserDb *UserDb) Add(c *gin.Context, User *model.User) error {
	rst := UserDb.db.Create(User)
	if ProcErrorLog(c, rst, User) {
		return rst.Error
	}
	return nil
}

func (UserDb *UserDb) Update(c *gin.Context, User *model.User) error {
	rst := UserDb.db.Updates(User)
	if ProcErrorLog(c, rst, User) {
		return nil
	}
	return rst.Error
}

func (UserDb *UserDb) One(c *gin.Context, UserId int64) (*model.User, error) {
	User := new(model.User)
	rst := UserDb.db.First(User, UserId)
	if ProcErrorLog(c, rst, User) {
		return nil, rst.Error
	}
	return User, nil
}

func (UserDb *UserDb) List(c *gin.Context, query *model.User, page, size int) ([]*model.User, error) {
	Users := make([]*model.User, 0)
	rst := UserDb.db.Limit(size).Offset((page - 1) * size).Where(query).Find(&Users)
	if ProcErrorLog(c, rst, map[string]interface{}{"query": query, "page": page, "size": size}) {
		return nil, rst.Error
	}
	return Users, nil
}

func (UserDb *UserDb) Delete(c *gin.Context, UserId int64) error {
	User := new(model.User)
	User.Uid = UserId
	User.Status = -1
	User.UpdateTime = time.Now().Unix()
	rst := UserDb.db.Updates(User)
	if ProcErrorLog(c, rst, map[string]interface{}{"UserId": UserId}) {
		return rst.Error
	}
	return nil
}
