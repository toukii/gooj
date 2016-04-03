package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/everfore/exc"
	"github.com/shaalx/gooj"
	"github.com/shaalx/gooj/goojle/models"
	"github.com/shaalx/goutils"
	"strings"
)

type PuzzleController struct {
	SessionController
}

func (c *PuzzleController) Prepare() {
	c.SessionController.Prepare()
}

// @router /puzzle [get]
func (c *PuzzleController) Puzzle_New() {
	c.TplName = "puzzle.html"
}

// @router /puzzle [post]
func (c *PuzzleController) PuzzlePost_New() {
	var puzzle models.Puzzle
	err := c.ParseForm(&puzzle)
	puzzle.Id = 0
	if !goutils.CheckErr(err) {
		puzzle.User = c.CurUser()
		n, err := models.ORM.Insert(&puzzle)
		if !goutils.CheckErr(err) {
			c.Redirect(fmt.Sprintf("/problem/%d", n), 302)
		}
		beego.Debug(n, err)
	}
	fmt.Println(puzzle)
}

// @router /puzzle/:id:int [get]
func (c *PuzzleController) Puzzle() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	var puzzle models.Puzzle
	err := models.ORM.QueryTable((*models.Puzzle)(nil)).Filter("Id", id).One(&puzzle)
	goutils.CheckErr(err)
	c.Data["puzzle"] = puzzle
	c.TplName = "puzzle.html"
}

// @router /puzzle/:id:int [post]
func (c *PuzzleController) PuzzlePostId() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	var puzzle models.Puzzle
	err := c.ParseForm(&puzzle)
	puzzle.Id = id
	if !goutils.CheckErr(err) {
		puzzle.User = c.CurUser()
		n, err := models.ORM.Update(&puzzle)
		beego.Debug(n, err)
	}
	fmt.Println(puzzle)
	c.Redirect(fmt.Sprintf("/problem/%d", id), 302)
}

// @router /test [post]
func (c *PuzzleController) Test() {
	submit_LOCKER.Lock()
	defer submit_LOCKER.Unlock()
	var model gooj.Model
	model.Id = "1"
	model.Desc = c.GetString("descr")
	model.Title = c.GetString("title")
	model.FuncName = c.GetString("func_name")
	model.ArgsType = c.GetString("args_type")
	model.Content = c.GetString("content")
	model.RetsType = c.GetString("rets_type")
	model.TestCases = c.GetString("test_cases")

	fmt.Printf("%#v\n", model)
	path_ := strings.Split(c.Ctx.Request.RemoteAddr, ":")[1]
	if len(path_) <= 1 {
		path_ = "goojt"
	}
	beego.Debug("path_:", path_)
	err := gooj.GenerateOjModle(path_, &model)
	if goutils.CheckErr(err) {
		return
	}
	cmd := exc.NewCMD("go test -v")
	ret, err := cmd.Wd().Cd(path_).Debug().Do()
	if goutils.CheckErr(err) {
		return
	}
	c.Ctx.ResponseWriter.Write(ret)
	fmt.Println(goutils.ToString(ret))
	go cmd.Reset(fmt.Sprintf("rm -rf %s", path_)).Cd(defaultpath).ExecuteAfter(1)
}
