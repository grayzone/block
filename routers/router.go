package routers

import (
	"github.com/astaxie/beego"
	"github.com/grayzone/block/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/new", &controllers.MainController{}, "GET:New")
}
