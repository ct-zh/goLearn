package repositories

import (
	"database/sql"
	"go_iris/datamodels"
)

type IOrder interface {
	Conn() error
	Insert(*datamodels.Order) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Order) error
	SelectByKey(int64) (*datamodels.Order, error)
	SelectAll() ([]*datamodels.Order, error)
}

type OrderManager struct {
	table  string
	dbConn *sql.DB
}

func (o OrderManager) Conn() error {
	panic("implement me")
}

func (o OrderManager) Insert(*datamodels.Order) (int64, error) {
	panic("implement me")
}

func (o OrderManager) Delete(int64) bool {
	panic("implement me")
}

func (o OrderManager) Update(*datamodels.Order) error {
	panic("implement me")
}

func (o OrderManager) SelectByKey(int64) (*datamodels.Order, error) {
	panic("implement me")
}

func (o OrderManager) SelectAll() ([]*datamodels.Order, error) {
	panic("implement me")
}

func NewOrderManager(t string, mysql *sql.DB) IOrder {
	return &OrderManager{
		table:  t,
		dbConn: mysql,
	}
}
