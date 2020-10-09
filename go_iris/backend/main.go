package backend

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
	template := iris.HTML("./backend/web/views", ".html").Layout(
		"shared/layout.html").Reload(
		true)
	app.RegisterView(template)

	// 4.  设置模版
	//app.StaticWeb("/assets", "./backend/web/assets")

	// 5. 异常跳转
	app.OnAnyErrorCode(func(context iris.Context) {
		context.ViewData("message",
			context.Values().GetStringDefault("message", "访问页面出错"))
		context.ViewLayout("")
		context.View("shared/error.html")
	})

	// 6. 连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 注册控制器
	//mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	productRepository := repositories.NewProductManage("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr("localhost:12999"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
