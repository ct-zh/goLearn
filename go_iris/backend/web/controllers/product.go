package controllers

import (
	"github.com/kataras/iris/v12"
	"go_iris/datamodels"
	"go_iris/services"
	"go_iris/tool"
	"strconv"
)

type ProductController struct {
	Ctx     iris.Context
	Service services.IProductService
}

func (p *ProductController) GetAll() interface{} {
	products, _ := p.Service.GetAll()
	return tool.Output{
		Code: 200,
		Msg:  "",
		Data: products,
	}
}

func (p *ProductController) PostUpdate() interface{} {
	id, err := strconv.ParseInt(p.Ctx.PostValue("id"), 10, 64)
	if err != nil {
		return tool.Output{
			Code: -1,
			Msg:  "id非法",
		}
	}

	product, err := p.Service.GetById(id)
	if err != nil {
		return tool.Output{
			Code: -1,
			Msg:  err.Error(),
		}
	}

	var (
		productName  = p.Ctx.PostValue("productName")
		productImage = p.Ctx.PostValue("productImage")
		productUrl   = p.Ctx.PostValue("productUrl")
	)

	flag := false

	productNum, err := strconv.ParseInt(p.Ctx.PostValue("productNum"), 10, 64)
	if err == nil && productNum > 0 {
		product.ProductNum = productNum
		flag = true
	}
	if len(productName) > 0 {
		product.ProductName = productName
		flag = true
	}
	if len(productImage) > 0 {
		product.ProductImage = productImage
		flag = true
	}
	if len(productUrl) > 0 {
		product.ProductUrl = productUrl
		flag = true
	}

	if !flag {
		return tool.Output{
			Code: 0,
			Msg:  "没有需要更新的数据",
			Data: product,
		}
	}

	if err := p.Service.Update(product); err != nil {
		return tool.Output{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return tool.Output{
		Code: 0,
		Msg:  "success",
		Data: product,
	}
}

func (p *ProductController) PostAdd() interface{} {
	var (
		productName  = p.Ctx.PostValue("productName")
		productImage = p.Ctx.PostValue("productImage")
		productUrl   = p.Ctx.PostValue("productUrl")
	)
	productNum, err := strconv.ParseInt(p.Ctx.PostValue("productNum"), 10, 64)
	if err != nil || productNum <= 0 {
		return tool.Output{
			Code: -1,
			Msg:  "商品库存非法",
		}
	}

	if len(productName) == 0 {
		return tool.Output{
			Code: -1,
			Msg:  "商品标题不能为空",
		}
	}

	if len(productImage) == 0 {
		return tool.Output{
			Code: -1,
			Msg:  "商品图片不能为空",
		}
	}

	if len(productUrl) == 0 {
		return tool.Output{
			Code: -1,
			Msg:  "商品地址不能为空",
		}
	}

	product := datamodels.Product{
		ProductName:  productName,
		ProductNum:   productNum,
		ProductImage: productImage,
		ProductUrl:   productUrl,
	}

	_, err = p.Service.Insert(&product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
		return tool.Output{
			Code: -1,
			Msg:  err.Error(),
		}
	}

	return tool.Output{
		Code: 0,
		Msg:  "success",
		Data: product,
	}
}

func (p *ProductController) GetOne() interface{} {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
		return tool.Output{
			Code: -1,
			Msg:  err.Error(),
		}
	}

	product, err := p.Service.GetById(id)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
		return tool.Output{
			Code: -1,
			Msg:  err.Error(),
		}
	}

	return tool.Output{
		Code: 0,
		Msg:  "success",
		Data: product,
	}
}

func (p *ProductController) GetDelete() interface{} {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	if p.Service.DeleteById(id) {
		return tool.Output{
			Code: -1,
			Msg:  "删除商品成功，ID为：" + idString,
		}
	} else {
		return tool.Output{
			Code: -1,
			Msg:  "删除商品失败，商品可能已被删除",
		}
	}
}
