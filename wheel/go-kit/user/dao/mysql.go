package dao

import (
	"fmt"
	"log"

	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// 初始化mysql连接
func InitMysql(dbConf conf.Db) (err error) {
	db, err = gorm.Open("mysql",
		// user:password@tcp(host:port)/dbName?charset=utf8&parseTime=True&loc=Local
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
			dbConf.User, dbConf.Passwd, dbConf.Host, dbConf.Port, dbConf.DbName))
	if err != nil {
		log.Println(err)
		return
	}

	// 使用单数表名
	// gorm默认的结构体映射是复数形式，比如你的博客表为blog，对应的结构体名就会是blogs，
	// 同时若表名为多个单词，对应的model结构体名字必须是驼峰式，首字母也必须大写
	// 如不禁用则会出现 表名为y结尾 的表 y变成ies的问题
	// 如果只是部分表需要使用源表名，请在实体类中声明TableName的构造函数
	// ```
	// func (实体名) TableName() string {
	// 		return "数据库表名"
	// }
	// ```
	db.SingularTable(true)
	return
}
