package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/shaalx/gooj"
	"github.com/shaalx/gooj/goojle/models"
	// "github.com/shaalx/goutils"
)

type JudgeController struct {
	SessionController
}

// @router / [get]
func (c *JudgeController) Get() {
	var puzzles []models.Puzzle
	n, err := models.ORM.QueryTable((*models.Puzzle)(nil)).RelatedSel().Filter("Online", 1).Limit(20).All(&puzzles)
	beego.Debug(n, err)
	c.Data["title"] = "Puzzle"
	puzzlez := make([]models.Puzzle, 0, len(puzzles))
	for _, it := range puzzles {
		(&it).SubString(100)
		puzzlez = append(puzzlez, it)
	}
	// fmt.Println(puzzlez)
	c.Data["puzzles"] = puzzlez
	c.TplName = "list.html"
}

// @router /state [get]
func (c *JudgeController) State() {
	var solutions []models.Solution
	n, err := models.ORM.QueryTable((*models.Solution)(nil)).RelatedSel().OrderBy("-created").Limit(12).All(&solutions)
	beego.Debug(n, err)
	c.Data["title"] = "State"
	c.Data["solutions"] = solutions
	c.TplName = "state.html"
}

func Puzzle2Model(puzzle *models.Puzzle) *gooj.Model {
	if nil == puzzle {
		return nil
	}
	m := gooj.Model{}
	m.Id = fmt.Sprintf("%s", puzzle.Id)
	m.ArgsType = puzzle.ArgsType
	m.Content = puzzle.Content
	m.Desc = puzzle.Descr
	m.FuncName = puzzle.FuncName
	m.Online = puzzle.Online
	m.RetsType = puzzle.RetsType
	m.TestCases = puzzle.TestCases
	m.Title = puzzle.Title
	return &m
}
