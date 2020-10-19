package controllers

import (
	"github.com/kataras/iris/v12"
	"go_iris/services"
	"go_iris/tool"
)

type OrderController struct {
	Ctx     iris.Context
	Service services.IOrderService
}

func (o *OrderController) Get() interface{} {
	orderArr, err := o.Service.GetAll()
	if err != nil {
		return tool.Output{
			Code: -1,
			Msg:  err.Error(),
		}
	}

	return tool.Output{
		Code: 0,
		Msg:  "success",
		Data: orderArr,
	}
}
