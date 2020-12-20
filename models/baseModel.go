package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"spider.ohmgood.com/common"
)

type BaseModel struct {
}

func init() {
	// set default database
	host, err := config.String("DB_HOST")
	if err != nil {
		common.Logs.Warning("DB_HOST_EMPTY", err.Error())
	}
	userName, err := config.String("DB_USER")
	if err != nil {
		common.Logs.Warning("DB_USER_EMPTY", err.Error())
	}
	pass, err := config.String("DB_PASSWORD")
	if err != nil {
		common.Logs.Warning("DB_PASSWORD_EMPTY", err.Error())
	}
	dbName, err := config.String("DB_NAME")
	if err != nil {
		common.Logs.Warning("DB_NAME_EMPTY", err.Error())
	}
	port, err := config.String("DB_PORT")
	if err != nil {
		common.Logs.Warning("DB_PORT_EMPTY", err.Error())
	}
	charset, err := config.String("DB_CHARSET")
	if err != nil {
		common.Logs.Warning("DB_CHARSET_EMPTY", err.Error())
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", userName+":"+pass+"@tcp("+host+":"+port+")/"+dbName+"?charset="+charset+"&loc=Local")

	// register model
	orm.RegisterModel(new(Job))
}
