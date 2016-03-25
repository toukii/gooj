package routers

import (
	"github.com/astaxie/beego"
	"github.com/shaalx/gooj/goojle/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.LogController{}, &controllers.MainController{}, &controllers.RegisterController{})
	beego.ErrorController(&controllers.ErrorController{})
}
