package routers

import (
	"LoginUser/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.LoginController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logged", &controllers.LoginController{}, "get:Logged")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
}
