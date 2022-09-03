package database

import (
	"encoding/json"
	"fmt"
	rawlog "log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	"github.com/gin-gonic/gin"
	"github.com/xilylg/accountbook/log"
	"github.com/xilylg/accountbook/utils"
)

type DatabaseHostConf struct {
	Host         string
	Port         int32
	Username     string
	Password     string
	Databasename string
}
type AccountbookDb struct {
	Master []*DatabaseHostConf
	Slave  []*DatabaseHostConf
}
type DatabaseConf struct {
	Driver      string
	Pretable    string
	Logfile     string
	Accountbook *AccountbookDb
}

type DbServiceSrt struct {
	Conf *DatabaseConf
	Db   *gorm.DB
}

var DbService *DbServiceSrt

func Init(dbConf *DatabaseConf) (*DbServiceSrt, error) {
	DbService = new(DbServiceSrt)
	DbService.Conf = dbConf
	db, err := conn(dbConf)
	if err != nil {
		return nil, err
	}
	DbService.Db = db

	//初始化各表服务
	Member = registerMember(db)
	Book = registerBook(db)
	Item = registerItem(db)
	Target = registerTarget(db)
	Flow = registerFlow(db)
	User = registerUser(db)

	return DbService, nil
}

func conn(conf *DatabaseConf) (*gorm.DB, error) {
	mDsn0 := ""
	var Sources, Replicas []gorm.Dialector
	switch conf.Driver {
	case "mysql":
		mDsn0, Sources, Replicas = MysqlConn(conf)
	}
	file, _ := os.OpenFile(conf.Logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	newLogger := logger.New(
		rawlog.New(file, "\r\n", rawlog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	preTable := ""
	if conf.Pretable != "" {
		preTable = conf.Pretable + "_"
	}

	db, err := gorm.Open(mysql.Open(mDsn0), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   preTable,
			SingularTable: true,
		},
	})
	db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  Sources,
		Replicas: Replicas,
		Policy:   dbresolver.RandomPolicy{},
	}).SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(100).
		SetMaxOpenConns(200))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MysqlConn(conf *DatabaseConf) (string, []gorm.Dialector, []gorm.Dialector) {
	connStr := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	mDsn0 := ""

	Sources := make([]gorm.Dialector, 0)
	for i, mConf := range conf.Accountbook.Master {
		dsn := fmt.Sprintf(connStr,
			mConf.Username, mConf.Password, mConf.Host, mConf.Port, mConf.Databasename)
		if i == 0 {
			mDsn0 = dsn
		}
		Sources = append(Sources, mysql.Open(dsn))
	}

	Replicas := make([]gorm.Dialector, 0)
	for _, sConf := range conf.Accountbook.Slave {
		dsn := fmt.Sprintf(connStr,
			sConf.Username, sConf.Password, sConf.Host, sConf.Port, sConf.Databasename)
		Replicas = append(Replicas, mysql.Open(dsn))
	}

	return mDsn0, Sources, Replicas
}

func ProcErrorLog(c *gin.Context, rst *gorm.DB, valObj interface{}) bool {
	if rst.Error == nil {
		return false
	}
	val, err := json.Marshal(valObj)
	if err != nil {
		log.Error(c, "database.json", rst.Error.Error())
		return true
	}
	log.Error(c, "database", rst.Error.Error()+" value:"+string(val)+utils.GetCallerInfo(2))
	return true
}
