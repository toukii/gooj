package routers

import (
	"github.com/astaxie/beego"
	"github.com/shaalx/gooj/goojle/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.MainController{}, &controllers.RegisterController{}, &controllers.LogController{})
	beego.ErrorController(&controllers.ErrorController{})
}
