package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"go_iris/datamodels"
	"go_iris/encrypt"
	"go_iris/services"
	"go_iris/tool"
	"strconv"
)

type UserController struct {
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

func (c *UserController) GetRegister() mvc.View {
	return mvc.View{
		Name: "user/register.html",
	}
}

func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "user/login.html",
	}
}

func (c *UserController) PostRegister() {
	var (
		nickName = c.Ctx.PostValue("nickName")
		userName = c.Ctx.PostValue("userName")
		password = c.Ctx.PostValue("password")
	)

	c.Ctx.Application().Logger().Debug(
		fmt.Sprintf("post params: Nick: %s, Account: %s, Pwd: %s",
			nickName,
			userName,
			password))

	// wait for params validate
	// ozzo-validation
	user := &datamodels.User{
		NickName: nickName,
		Account:  userName,
		Password: password,
	}

	_, err := c.Service.AddUser(user)
	if err != nil {
		c.Ctx.Redirect("/user/error")
		return
	}

	c.Ctx.Redirect("/user/login")
	return
}

func (c *UserController) PostLogin() mvc.Response {
	var (
		account  = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)

	user, isOk := c.Service.IsPwdSuccess(account, password)
	if !isOk {
		return mvc.Response{
			Path: "/user/login",
		}
	}

	tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))

	uidByte := []byte(strconv.FormatInt(user.ID, 10))
	uidString, err := encrypt.EnPwdCode(uidByte)
	if err != nil {
		c.Ctx.Application().Logger().Debug(err)
	}
	tool.GlobalCookie(c.Ctx, "uid", uidString)

	return mvc.Response{
		Path: "/product/",
	}
}
