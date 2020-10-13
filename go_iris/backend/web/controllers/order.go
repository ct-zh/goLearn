package controllers

import (
	"github.com/kataras/iris/v12"
	"go_iris/services"
)

type OrderController struct {
	Ctx     iris.Context
	Service services.IOrderService
}
