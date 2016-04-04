package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:JudgeController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:JudgeController"],
		beego.ControllerComments{
			"Get",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:JudgeController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:JudgeController"],
		beego.ControllerComments{
			"State",
			`/state`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:ListController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:ListController"],
		beego.ControllerComments{
			"GetPro",
			`/oj/:id:int`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:ListController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:ListController"],
		beego.ControllerComments{
			"OJ",
			`/oj/:id:int`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LoginController"],
		beego.ControllerComments{
			"LoadLogin",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LoginController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:LoginController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
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

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegistController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegistController"],
		beego.ControllerComments{
			"Signin",
			`/githubsignin`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegistController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegistController"],
		beego.ControllerComments{
			"Callback",
			`/callback`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegistController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegistController"],
		beego.ControllerComments{
			"LoadRegister",
			`/register`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegistController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:RegistController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:UserController"],
		beego.ControllerComments{
			"User",
			`/user`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:UserController"],
		beego.ControllerComments{
			"Puzzles",
			`/puzzles`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shaalx/gooj/goojle/controllers:UserController"],
		beego.ControllerComments{
			"Solutions",
			`/solutions`,
			[]string{"get"},
			nil})

}
