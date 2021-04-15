package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetTestDb() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	return db, nil
}
