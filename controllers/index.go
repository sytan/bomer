// Package controllers implement router handler
package controllers

import (
	"github.com/astaxie/beego"
)

const (
	indexTPL     = "index.tpl"
	loginTPL     = "login.tpl"
	dashboardTPL = "dashboard.tpl"
)

// IndexController customerize beego controller
type IndexController struct {
	beego.Controller
}

// Get implement http get
func (c *IndexController) Get() {
	c.Data["website"] = "http://www.invoxia.com"
	c.Data["company"] = "invoxia"
	c.Data["email"] = "sy.tang@invoxia.com"
	c.Data["author"] = "Sytang"
	c.TplName = indexTPL
}
