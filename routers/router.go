package routers

import (
	"github.com/astaxie/beego"
	"github.com/grayzone/block/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/new", &controllers.NewController{})
	beego.Router("/next", &controllers.NextController{})

	beego.Router("/donew", &controllers.NewController{}, "GET:Do")
	beego.Router("/donext", &controllers.NextController{}, "GET:Do")
}
