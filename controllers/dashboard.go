// Package controllers implement router handler
package controllers

import "github.com/astaxie/beego"

// DashboardController implement controller for login
type DashboardController struct {
	beego.Controller
}

type user struct {
	UserName string
	Password string
}

// UserOnline record information of user login
var UserOnline user

// Get implement http Get
func (c *DashboardController) Get() {
	c.TplName = dashboardTPL
}

// Post implement http post
func (c *DashboardController) Post() {
	UserOnline.UserName = c.GetString("email")
	UserOnline.Password = c.GetString("password")
	if UserOnline.UserName == "tashy1@163.com" && UserOnline.Password == "swissvoice" {
		c.Data["username"] = UserOnline.UserName
		c.TplName = dashboardTPL
	} else {
		c.Data["failLogin"] = "Invalid user name and password, Please try again!"
		c.Redirect("/", 302) //don't understand why it's 302
	}
}
