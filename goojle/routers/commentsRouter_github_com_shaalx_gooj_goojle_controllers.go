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
			"LoadRegister",
			`/register`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LogController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
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

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"],
		beego.ControllerComments{
			"LoadLogin",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:MainController"],
		beego.ControllerComments{
			"User",
			`/user`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"],
		beego.ControllerComments{
			"Puzzle_New",
			`/puzzle`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"],
		beego.ControllerComments{
			"PuzzlePost_New",
			`/puzzle`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"],
		beego.ControllerComments{
			"Puzzle",
			`/puzzle/:id:int`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"],
		beego.ControllerComments{
			"PuzzlePostId",
			`/puzzle/:id:int`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:PuzzleController"],
		beego.ControllerComments{
			"Test",
			`/test`,
			[]string{"post"},
			nil})

}
