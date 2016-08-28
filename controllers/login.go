// Package controllers implement router handler
package controllers

import "github.com/astaxie/beego"

// LoginController implement controller for login
type LoginController struct {
	beego.Controller
}

// Get implement http Get
func (c *LoginController) Get() {
	c.TplName = loginTPL
}

// Post implement http post
func (c *LoginController) Post() {
	username := c.GetString("email")
	password := c.GetString("password")
	if username == "tashy1@163.com" && password == "swissvoice" {
		c.TplName = dashboardTPL
	} else {
		c.Data["failLogin"] = "Invalid user name and password, Please try again!"
		c.TplName = loginTPL
	}
}
