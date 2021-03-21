package middleware

import "github.com/kataras/iris/v12"

func AuthConProduct(ctx iris.Context) {
	uid := ctx.GetCookie("uid")
	if uid == "" {
		ctx.Application().Logger().Debug("未登陆")
		ctx.Redirect("/user/login")
		return
	} else {
		ctx.Application().Logger().Debug("未登陆")
		ctx.Next()
	}
}
