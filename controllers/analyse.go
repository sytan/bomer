// Package controllers implement router handler
package controllers

import (
	"github.com/astaxie/beego"
)

// AnalyseController implement controller for login
type AnalyseController struct {
	beego.Controller
}

// Get implement http get
func (c *AnalyseController) Get() {
	c.TplName = analyseTPL
}

// Post implement http post
func (c *AnalyseController) Post() {
	c.Data["startIndex"] = c.GetString("startIndex")
	c.TplName = analyseTPL
}
