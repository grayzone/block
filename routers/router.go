package routers

import (
	"github.com/astaxie/beego"
	"github.com/grayzone/block/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/new", &controllers.MainController{}, "GET:New")
	beego.Router("/test", &controllers.MainController{}, "GET:Test")
	beego.Router("/drop", &controllers.MainController{}, "POST:Drop")

}
