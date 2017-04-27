package controllers

import (
	"github.com/astaxie/beego"
)

type NextController struct {
	beego.Controller
}

func (c *NextController) Get() {
	c.TplName = "next.html"
	c.Layout = "layout.html"
}

func (c *NextController) Do() {
	result := "/next"
	c.Data["json"] = &result
	c.ServeJSON()
}
