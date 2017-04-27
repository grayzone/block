package controllers

import (
	"github.com/astaxie/beego"
)

type NewController struct {
	beego.Controller
}

func (c *NewController) Get() {
	c.TplName = "new.html"
	c.Layout = "layout.html"
}

func (c *NewController) Do() {
	result := "/new"
	c.Data["json"] = &result
	c.ServeJSON()
}
