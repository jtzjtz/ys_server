package sqldb

import (
	"github.com/jtzjtz/kit/conn/mysql_pool"
	"github.com/jtzjtz/ys_pack/server/dao"
	"github.com/jtzjtz/ys_server/app/config"
	"gorm.io/gorm"
)

func Init() {
	dao.SqlDBWrite = NewSqlDatabase(false)
	dao.SqlDBRead = NewSqlDatabase(true)
}

func NewSqlDatabase(isRead bool) *gorm.DB {

	var userName, password, databaseName, host, port string

	if isRead {
		userName = config.SqlConfig.DbUser
		password = config.SqlConfig.DbPwd
		databaseName = config.SqlConfig.DbName
		host = config.SqlConfig.DbHost
		port = config.SqlConfig.DbPort
	} else {
		userName = config.SqlConfig.DbUser
		password = config.SqlConfig.DbPwd
		databaseName = config.SqlConfig.DbName
		host = config.SqlConfig.DbHost
		port = config.SqlConfig.DbPort
	}

	options := &mysql_pool.Options{
		InitCap:  50,
		MaxCap:   100,
		IsDebug:  true,
		User:     userName,
		Pass:     password,
		Host:     host,
		Port:     port,
		DataBase: databaseName,
	}

	DB, _ := mysql_pool.NewMySqlPool(options)

	return DB

}
