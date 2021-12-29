package config

import (
	"github.com/jtzjtz/kit/config"
)

var (
	client    config.Client
	SqlConfig mysqlConfig
)

func Init(e string) (err error) {
	SqlConfig.DbHost = "localhost"
	SqlConfig.DbName = "hys_db"
	SqlConfig.DbPort = "3306"
	SqlConfig.DbPwd = "123456"
	SqlConfig.DbUser = "root"
	//nacasConfig := config.NacosConfig{
	//	Address:             "10.31.114.24",
	//	Port:                8848,
	//	Scheme:              "http",
	//	ContextPath:         "/nacos",
	//	NameSpaceId:         "fffaacb2-a7f7-45d5-83aa-dec49122a098",
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "./.config/log",
	//	CacheDir:            "./.config",
	//	RotateTime:          "1h",
	//	MaxAge:              3,
	//	LogLevel:            "debug",
	//}
	//client, err = config.InitClient(nacasConfig)
	//if err != nil {
	//	return
	//}
	//client.GetDataConfig("mysql_content", "grpc_service", &SqlConfig, true)
	return
}
