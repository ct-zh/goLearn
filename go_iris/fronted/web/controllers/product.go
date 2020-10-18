package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"go_iris/datamodels"
	"go_iris/services"
	"strconv"
)

type ProductController struct {
	Ctx     iris.Context
	Service services.IProductService
	Order   services.IOrderService
	Session *sessions.Session
}

func (p *ProductController) GetDetail() mvc.View {
	product, err := p.Service.GetById(1)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/view.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetOrder() mvc.View {
	var (
		id  = p.Ctx.URLParam("id")
		uid = p.Ctx.GetCookie("uid")
	)
	if uid == "" {
		p.Ctx.Application().Logger().Error("请重新登陆")
		p.Ctx.Redirect("user/login.html")
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}

	product, err := p.Service.GetById(int64(productId))
	if err != nil {
		p.Ctx.Application().Logger().Error(err)
	}

	var orderId int64
	showMsg := "抢购失败"
	if product.ProductNum > 0 {
		product.ProductNum -= 1
		err := p.Service.Update(product)
		if err != nil {
			p.Ctx.Application().Logger().Error(err)
		}

		userId, err := strconv.Atoi(uid)
		if err != nil {
			p.Ctx.Application().Logger().Error(err)
		}

		// 创建订单
		order := &datamodels.Order{
			UserId:      int64(userId),
			ProductId:   int64(product.ID),
			OrderStatus: datamodels.OrderSuccess,
		}

		orderId, err = p.Order.Insert(order)
		if err != nil {
			p.Ctx.Application().Logger().Error(err)
		} else {
			showMsg = "抢购成功"
		}
	}

	return mvc.View{
		Layout: "shared/productLayout.html",
		Name:   "product/result.html",
		Data: iris.Map{
			"orderID":     orderId,
			"showMessage": showMsg,
		},
	}
}
