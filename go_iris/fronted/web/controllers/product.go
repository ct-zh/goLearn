package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"go_iris/datamodels"
	"go_iris/services"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
)

type ProductController struct {
	Ctx     iris.Context
	Service services.IProductService
	Order   services.IOrderService
	Session *sessions.Session
}

var (
	htmlOutPath  = "./fronted/web/htmlProductShow/"
	templatePath = "./fronted/web/views/template/"
)

func (p *ProductController) GetGenerateHtml() {
	id := p.Ctx.URLParam("id")
	productId, err := strconv.Atoi(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	// 获取模版文件地址
	contentTmp, err := template.ParseFiles(filepath.Join(templatePath), "product.html")
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	// 获取html生成路径
	fileName := filepath.Join(htmlOutPath, "htmlProduct.html")

	// 获取模版渲染数据
	// todo： 验证每次访问时会不会请求数据库
	product, err := p.Service.GetById(int64(productId))
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	// 生成静态文件
	generateStaticHtml(p.Ctx, contentTmp, fileName, product)
}

func generateStaticHtml(
	ctx iris.Context,
	template *template.Template,
	fileName string, product *datamodels.Product) {

	// 判断文件是否存在
	if exist(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			ctx.Application().Logger().Error(err)
		}
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
	defer file.Close()

	template.Execute(file, &product)

}

func exist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
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
