package repository

import (
	"github.com/jinzhu/gorm"
	"hot/domain/model"
)

type IAccount interface {
	// 主键查询
	Get(id int) (*model.FullTableScanTest, error)
	// 获取账号信息
	GetByAccount(string) (*model.FullTableScanTest, error)
	// 更新类型
	UpdateTypeById(test *model.FullTableScanTest) error
}

type Account struct {
	db *gorm.DB
}

func (a *Account) UpdateTypeById(test *model.FullTableScanTest) error {
	return a.db.Model(test).Update(test).Error
}

func (a *Account) Get(id int) (*model.FullTableScanTest, error) {
	acc := &model.FullTableScanTest{}
	return acc, a.db.Where("id = ?", id).First(acc).Error
}

func (a *Account) GetByAccount(s string) (*model.FullTableScanTest, error) {
	acc := &model.FullTableScanTest{}
	return acc, a.db.Where("account = ?", s).First(acc).Error
}

func NewAccount(db *gorm.DB) IAccount {
	return &Account{db: db}
}
