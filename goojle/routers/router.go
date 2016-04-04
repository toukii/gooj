package routers

import (
	"github.com/astaxie/beego"
	"github.com/shaalx/gooj/goojle/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.JudgeController{}, &controllers.LoginController{}, &controllers.RegistController{}, &controllers.PuzzleController{})
	beego.ErrorController(&controllers.ErrorController{})
}
