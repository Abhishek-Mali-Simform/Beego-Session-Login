package controllers

import (
	"LoginUser/library"
	"LoginUser/models"
	"LoginUser/services"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller

	IsLogin  bool
	UserInfo *models.User
}

func (login *LoginController) Index() {
	login.IsLogin = login.GetSession("userinfo") != nil
	if login.IsLogin {
		login.Ctx.Redirect(302, "/logged")
	}
	login.TplName = "login.html"
}

func (login *LoginController) Login() {
	login.IsLogin = login.GetSession("userinfo") != nil
	if login.IsLogin {
		login.Ctx.Redirect(302, "/logged")
	}
	if !login.Ctx.Input.IsPost() {
		return
	}

	flash := beego.NewFlash()
	email := login.GetString("Email")
	password := login.GetString("Password")

	user, err := library.AuthenticateUser(email, password)
	if err != nil && user.Id < 1 {
		flash.Warning(err.Error())
		flash.Store(&login.Controller)
		return
	}
	flash.Success("Success logged in")
	flash.Store(&login.Controller)
	login.SetLogin(user)
	login.Redirect(login.URLFor("LoginController.Logged"), 303)
}

func (login *LoginController) SetLogin(user *models.User) {
	fmt.Println(user.Id)
	login.SetSession("userinfo", user.Id)
	login.UserInfo = user
}

func (login *LoginController) GetLogin() *models.User {
	user := &models.User{Id: login.GetSession("userinfo").(int64)}
	err := services.Read(user)
	services.ErrorCheck("Cannot Read User. Reason: ", err)
	login.UserInfo = user
	return user
}

func (login *LoginController) Logged() {
	login.IsLogin = login.GetSession("userinfo") != nil
	if !login.IsLogin {
		login.Ctx.Redirect(302, login.URLFor("LoginController.Index"))
		return
	}
	login.Data["UserInfo"] = login.GetLogin()
	login.TplName = "index.tpl"
}

func (login *LoginController) Logout() {
	login.DelSession("userinfo")
	flash := beego.NewFlash()
	flash.Success("Success Logout")
	flash.Store(&login.Controller)
	login.Ctx.Redirect(302, login.URLFor("LoginController.Index"))
}
