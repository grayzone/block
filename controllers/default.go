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
