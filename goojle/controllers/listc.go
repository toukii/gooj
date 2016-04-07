package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shaalx/gooj/goojle/models"
	// "github.com/shaalx/goutils"
)

type JudgeController struct {
	SessionController
}

// @router / [get]
func (c *JudgeController) Get() {
	var puzzles []models.Puzzle
	n, err := models.ORM.QueryTable((*models.Puzzle)(nil)).RelatedSel().Filter("Online", 1).Limit(15, 0).All(&puzzles)
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
	n, err := models.ORM.QueryTable((*models.Solution)(nil)).RelatedSel().OrderBy("-created").Limit(15).All(&solutions)
	beego.Debug(n, err)
	c.Data["title"] = "State"
	c.Data["solutions"] = solutions
	c.TplName = "state.html"
}
