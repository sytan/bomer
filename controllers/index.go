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

// Post implement http post
func (c *IndexController) Post() {
	username := c.GetString("email")
	password := c.GetString("password")
	if username == "tashy1@163.com" && password == "swissvoice" {
		c.Data["username"] = username
		c.TplName = dashboardTPL
	} else {
		c.Data["failLogin"] = "Invalid user name and password, Please try again!"
		c.TplName = indexTPL
	}
}
