package routers

import (
	"github.com/astaxie/beego"
	"github.com/sy/bomer/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/dashboard", &controllers.DashboardController{})
	beego.Router("/import", &controllers.ImportController{})
	beego.Router("/analyse", &controllers.AnalyseController{})
}
