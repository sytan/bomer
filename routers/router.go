package routers

import (
	"github.com/astaxie/beego"
	"github.com/sy/bomanager/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
}
