package common

import "strconv"

type MysqlConf struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetLocalMysqlConf() MysqlConf {
	return MysqlConf{
		Host:     "127.0.0.1",
		User:     "root",
		Pwd:      "root",
		Database: "testdb",
		Port:     3306,
	}
}

func GetMysqlDSN(mysqlInfo MysqlConf) string {
	return mysqlInfo.User + ":" + mysqlInfo.Pwd +
		"@tcp(" + mysqlInfo.Host + ":" + strconv.FormatInt(mysqlInfo.Port, 10) + ")/" +
		mysqlInfo.Database + "?charset=utf8&parseTime=True&loc=Local"
}
