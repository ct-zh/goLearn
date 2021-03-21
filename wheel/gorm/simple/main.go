package main

import (
	"fmt"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/gorm/common"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/gorm/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", common.GetMysqlDSN(common.GetLocalMysqlConf()))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//testPro := &model.TestProduct{}
	//if !db.HasTable(testPro) {
	//	if err := db.CreateTable(testPro).Error; err != nil {
	//		panic(err)
	//	}
	//}

	// gorm简单的curd
	curd(db)
}

func curd(db *gorm.DB) {
	// create or insert
	newItem := &model.TestProduct{
		Code:  "AAbb",
		Price: 100,
	}
	if err := db.Create(newItem).Error; err != nil {
		panic(err)
	} else {
		fmt.Println("new Item 新增成功！")
	}

	newItem2 := &model.TestProduct{
		Code:  "ccDD",
		Price: 10,
	}
	if err := db.Create(newItem2).Error; err != nil {
		panic(err)
	} else {
		fmt.Println("new Item2 新增成功！")
	}

	// read or select
	var getItem model.TestProduct
	if err := db.First(&getItem, "code = ?", "AAbb").Error; err != nil {
		panic(err)
	} else {
		fmt.Printf("查询item成功： %+v \n", getItem)
	}

	// update
	if err := db.Model(&getItem).Update("Price", 102).Error; err != nil {
		panic(err)
	} else {
		fmt.Println("更新成功!")
	}

	// delete 并非真正删除，只是更新deleted_at
	if err := db.Delete(&getItem).Error; err != nil {
		panic(err)
	} else {
		fmt.Println("删除成功")
	}
}
