package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go_iris/backend/web/controllers"
	"go_iris/common"
	"go_iris/repositories"
	"go_iris/services"
)

func main() {
	// 1. 创建 iris 实例
	app := iris.New()

	// 2. 设置日志等级
	app.Logger().SetLevel("debug")

	// 3. 注册模版
	//template := iris.HTML("./backend/web/views", ".html").Layout(
	//	"shared/layout.html").Reload(
	//	true)
	//app.RegisterView(template)

	// 4.  设置模版
	// 旧版本的方法： app.StaticWeb("/assets", "./backend/web/assets")
	//app.HandleDir("/assets", iris.Dir("./backend/web/assets"))

	// 5. 异常跳转
	app.OnAnyErrorCode(func(context iris.Context) {
		context.JSON(iris.Map{"code": "-1", "msg": "404 NOT FOUND"})
		//context.ViewData("message", context.Values().GetStringDefault("message", "访问页面出错"))
		//context.ViewLayout("")
		//context.View("shared/error.html")
	})

	// 6. 连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 注册控制器
	mvc.New(app.Party("/product")). // 注册路径 /product
					Register( // 注册 IProductService 与 iris.context
			ctx,
								services.NewProductService(repositories.NewProductManage(db))).
		Handle(new(controllers.ProductController)) // 绑定到 productController

	mvc.New(app.Party("/order")).
		Register(
			ctx,
			services.NewService(repositories.NewOrderManager("order", db))).
		Handle(new(controllers.OrderController))

	app.Run(
		iris.Addr("localhost:12999"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
