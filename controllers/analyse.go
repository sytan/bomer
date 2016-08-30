// Package controllers implement router handler
package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// AnalyseController implement controller for login
type AnalyseController struct {
	beego.Controller
}

type formData struct {
	StartRow    string   `form:"startIndex"`
	TableHeader []string `form:"tableHeader"`
}

// Get implement http get
func (c *AnalyseController) Get() {
	c.TplName = analyseTPL
}

// Post implement http post
func (c *AnalyseController) Post() {
	analysisCondition := formData{}
	if err := c.ParseForm(&analysisCondition); err != nil {
		fmt.Println(err)
	} else {
		c.Data["startIndex"] = analysisCondition.StartRow
	}
	c.TplName = analyseTPL
}
