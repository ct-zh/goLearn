package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"go_iris/datamodels"
	"go_iris/services"
	"go_iris/tool"
	"strconv"
)

type UserController struct {
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

func (c *UserController) PostRegister() interface{} {
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
		return map[string]string{"code": "-1", "msg": err.Error()}
	}

	return map[string]string{"code": "0", "msg": "success"}
}

func (c *UserController) PostLogin() interface{} {
	var (
		account  = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)

	user, isOk := c.Service.IsPwdSuccess(account, password)
	if !isOk {
		return map[string]string{"code": "-1", "msg": "error password!"}
	}

	tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))
	c.Session.Set("userId", strconv.FormatInt(user.ID, 10))

	return map[string]string{"code": "0", "msg": "success!"}
}
