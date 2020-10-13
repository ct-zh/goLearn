package services

import "go_iris/repositories"

type IOrderService interface {
}

type OrderService struct {
	repository repositories.IOrder
}

func NewService(r repositories.IOrder) IOrderService {
	return &OrderService{repository: r}
}
