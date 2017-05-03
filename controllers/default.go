package controllers

import (
	"github.com/astaxie/beego"
	"github.com/grayzone/block/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	c.Layout = "layout.html"
}

func (c *MainController) New() {
	var result models.BlockBox
	result.Seed()
	c.Data["json"] = &result.Data
	c.ServeJSON()
}

func (c *MainController) Test() {
	var result models.BlockBox
	result.TestData()
	c.Data["json"] = &result.Data
	c.ServeJSON()
}

func (c *MainController) Drop() {
	s := c.GetString("data")
	var result models.BlockBox
	result.Parse(s)
	result.Format()

	c.Data["json"] = &result.Data
	c.ServeJSON()
}
