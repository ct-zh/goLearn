package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"go_iris/common"
	"go_iris/fronted/web/controllers"
	"go_iris/repositories"
	"go_iris/services"
	"time"
)

func main() {
	// 1. 创建iris实例
	app := iris.New()

	// 2. 设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")

	// 3. 注册模板
	temp := iris.HTML("./fronted/web/views",
		".html").Layout(
		"shared/layout.html").Reload(true)
	app.RegisterView(temp)

	// 4.  设置模版
	// 旧版本的方法： app.StaticWeb("/assets", "./fronted/web/assets")
	app.HandleDir("/assets", iris.Dir("./fronted/web/assets"))

	//访问生成好的html静态文件
	//app.HandleDir("/html", iris.Dir("./fronted/web/htmlProductShow"))

	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error/html")
	})

	// 6. 连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		panic(err)
	}

	sess := sessions.New(sessions.Config{
		Cookie:  "AdminCookie",
		Expires: 600 * time.Minute,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 注册控制器
	mvc.New(app.Party("/user")).Register(
		services.NewUserService(
			repositories.NewUserManage(db)),
		ctx, sess.Start,
	).Handle(new(controllers.UserController))

	mvc.New(app.Party("/product")).Register(
		services.NewProductService(
			repositories.NewProductManage(db)),
		ctx, sess.Start,
	).Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr(":12998"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
