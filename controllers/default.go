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
	result.TestData2()
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

func (c *MainController) Step() {
	s := c.GetString("data")
	var input models.BlockBox
	input.Parse(s)
	result := input.Step()

	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *MainController) Remove() {
	x, _ := c.GetInt("x")
	y, _ := c.GetInt("y")
	s := c.GetString("data")
	var result models.BlockBox
	result.Parse(s)
	result.Remove(x, y)
	c.Data["json"] = &result
	c.ServeJSON()
}
