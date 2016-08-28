// Package controllers implement router handler
package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// DashboardController implement controller for login
type DashboardController struct {
	beego.Controller
}

// FormController implement controller for form
// type FormController struct {
// 	beego.Controller
// }

// Get implement http Get
func (c *DashboardController) Get() {
	c.TplName = dashboardTPL
}

// Post implement http Post
func (c *DashboardController) Post() {
	f, h, err := c.GetFile("uploadFile")
	path := "./" + h.Filename
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		c.SaveToFile("uploadFile", path)
	}
	c.TplName = dashboardTPL
}

// Post implement http post
// func (c *LoginController) Post() {
// 	username := c.GetString("email")
// 	password := c.GetString("password")
// 	if username == "tashy1@163.com" && password == "swissvoice" {
// 		c.TplName = indexTPL
// 	} else {
// 		c.Data["failLogin"] = "Invalid user name and password, Please try again!"
// 		c.TplName = loginTPL
// 	}
// }
