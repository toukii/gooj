package routers

import (
	"github.com/astaxie/beego"
	"github.com/toukii/gooj/goojle/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.JudgeController{}, &controllers.LoginController{}, &controllers.RegistController{}, &controllers.PuzzleController{}, &controllers.UserController{}, &controllers.ListController{})
	beego.ErrorController(&controllers.ErrorController{})
	beego.SetStaticPath("public", "static")
}
