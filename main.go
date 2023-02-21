package main

import (
	_ "LoginUser/routers"
	"LoginUser/services"
	"LoginUser/utils"
	"github.com/astaxie/beego"
)

func init() {
	utils.LoadEnv()
	services.ConnectDb()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	//user := services.DummyUser()
	//services.Create(user)
	beego.Run()
}
