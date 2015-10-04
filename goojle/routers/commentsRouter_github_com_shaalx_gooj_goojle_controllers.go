package routers

import (
	"github.com/astaxie/beego"
)

func init() {

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

}
