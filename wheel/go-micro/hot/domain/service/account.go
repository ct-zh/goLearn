package service

import (
	"hot/domain/model"
	"hot/domain/repository"
)

type IAccount interface {
	// 获取账号信息
	GetByAccount(string) (*model.FullTableScanTest, error)
	// 更新类型
	UpdateTypeById(test *model.FullTableScanTest) error
}

type Account struct {
	repo repository.IAccount
}

func NewAccount(repo repository.IAccount) *Account {
	return &Account{repo: repo}
}

func (a *Account) GetByAccount(s string) (*model.FullTableScanTest, error) {
	return a.repo.GetByAccount(s)
}

func (a *Account) UpdateTypeById(test *model.FullTableScanTest) error {
	return a.repo.UpdateTypeById(test)
}
