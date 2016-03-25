package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"],
		beego.ControllerComments{
			"Signin",
			`/githubsignin`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"],
		beego.ControllerComments{
			"Callback",
			`/callback`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"],
		beego.ControllerComments{
			"LoadLogin",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"],
		beego.ControllerComments{
			"User",
			`/user`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"],
		beego.ControllerComments{
			"Update",
			`/update`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"],
		beego.ControllerComments{
			"GetPro",
			`/problem/:id:int`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"],
		beego.ControllerComments{
			"OJ",
			`/oj/:id:int`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegisterController"],
		beego.ControllerComments{
			"LoadRegister",
			`/register`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegisterController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

}
