package main

import (
	"context"
	"fmt"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/gorm/common"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(common.GetMysqlDSN(common.GetLocalMysqlConf())), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	select1(db)
}

// 新建会话模式测试
func select1(db *gorm.DB) {
	var p1 model.TestProduct
	tx := db.Session(&gorm.Session{DryRun: true}).Where("code = ?", "AAbb")

	res := tx.Where("price > ?", 1).Find(&p1)
	fmt.Printf("%+v \n %+v \n", res.Statement.SQL.String(), p1)

	var p2 model.TestProduct
	res2 := tx.Session(&gorm.Session{DryRun: true}).WithContext(context.TODO()).Where("price < ?", 10).Find(&p2)
	fmt.Printf("%+v \n %+v \n", res2.Statement.SQL.String(), p2)
}
