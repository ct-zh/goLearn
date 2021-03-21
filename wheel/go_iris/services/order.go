package services

import (
	"go_iris/datamodels"
	"go_iris/repositories"
)

type IOrderService interface {
	GetById(int64) (*datamodels.Order, error)
	DeleteById(int64) bool
	Update(order *datamodels.Order) error
	Insert(order *datamodels.Order) (int64, error)
	GetAll() ([]*datamodels.Order, error)
	GetAllWithInfo() (map[int]map[string]string, error)
	InsertByMsg(message *datamodels.Message) (int64, error)
}

type OrderService struct {
	repository repositories.IOrder
}

func (o *OrderService) GetById(id int64) (*datamodels.Order, error) {
	return o.repository.SelectByKey(id)
}

func (o *OrderService) DeleteById(id int64) bool {
	return o.repository.Delete(id)
}

func (o *OrderService) Update(order *datamodels.Order) error {
	return o.repository.Update(order)
}

func (o *OrderService) Insert(order *datamodels.Order) (int64, error) {
	return o.repository.Insert(order)
}

func (o *OrderService) GetAll() ([]*datamodels.Order, error) {
	return o.repository.SelectAll()
}

func (o *OrderService) GetAllWithInfo() (map[int]map[string]string, error) {
	return o.repository.SelectAllWithInfo()
}

func (o *OrderService) InsertByMsg(message *datamodels.Message) (int64, error) {
	order := &datamodels.Order{
		UserId:      message.UserId,
		ProductId:   message.ProductId,
		OrderStatus: datamodels.OrderSuccess,
	}
	return o.Insert(order)
}

func NewService(r repositories.IOrder) IOrderService {
	return &OrderService{repository: r}
}
