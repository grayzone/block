package controllers

import (
	"github.com/astaxie/beego"
	"github.com/grayzone/block/util"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
	c.Layout = "layout.html"
}

func (c *MainController) New() {
	result := util.GetSeedData()
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *MainController) Test() {
	result := [10][10]int{
		{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 2, 3, 4, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 3, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}}
	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *MainController) Drop() {
	s := c.GetString("data")
	drop := util.DropZeroData(s)
	left := util.LeftData(drop)

	c.Data["json"] = &left
	c.ServeJSON()
}
