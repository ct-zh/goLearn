package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go_iris/common"
	"go_iris/datamodels"
	"go_iris/services"
	"strconv"
)

type ProductController struct {
	Ctx     iris.Context
	Service services.IProductService
}

func (p *ProductController) GetAll() mvc.View {
	products, _ := p.Service.GetAll()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": products,
		},
	}
}

func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}

	p.parseForm(product)

	err := p.Service.Update(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

func (p *ProductController) PostAdd() {
	product := &datamodels.Product{}

	p.parseForm(product)

	_, err := p.Service.Insert(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) GetManager() mvc.View {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	product, err := p.Service.GetById(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	return mvc.View{
		Name: "product/manager.html",
		Data: iris.Map{
			"product": product,
		},
	}
}

func (p *ProductController) GetDelete() {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	if ok := p.Service.DeleteById(id); ok {
		p.Ctx.Application().Logger().Debug("删除商品成功，ID为：" + idString)
	} else {
		p.Ctx.Application().Logger().Debug("删除商品失败，ID为：" + idString)
	}
	p.Ctx.Redirect("/product/all")
}

func (p *ProductController) parseForm(product *datamodels.Product) {
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName: "imooc"})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
}
