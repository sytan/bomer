// Package controllers implement router handler
package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

// ImportController implement controller for login
type ImportController struct {
	beego.Controller
}

// Get implement http Get
func (c *ImportController) Get() {
	c.TplName = importTPL
}

// Post implement http Post
func (c *ImportController) Post() {
	f, h, err := c.GetFile("uploadFile")
	path := "./static/file/" + h.Filename
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		c.SaveToFile("uploadFile", path)
	}
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {

	}
	for _, sheet := range xlFile.Sheets {
		var maxLengthCells []*xlsx.Cell //to find the longest row
		for _, row := range sheet.Rows {
			if len(maxLengthCells) < len(row.Cells) {
				maxLengthCells = row.Cells //update to longer row
			}
		}
		c.Data["rows"] = sheet.Rows
		c.Data["maxLengthCells"] = maxLengthCells

	}
	c.Data["username"] = UserOnline.UserName
	c.TplName = importTPL
}
