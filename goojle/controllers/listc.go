package controllers

import (
	"github.com/astaxie/beego"
	"github.com/toukii/gooj/goojle/models"
	"github.com/toukii/gooj/goojle/utils"
	// "github.com/toukii/goutils"
)

type JudgeController struct {
	SessionController
}

// @router / [get]
func (c *JudgeController) Get() {
	page, _ := c.GetInt("page")
	var puzzles []models.Puzzle
	n, err := models.ORM.QueryTable((*models.Puzzle)(nil)).RelatedSel().Filter("Online", 1).Limit(15, 15*(page-1)).All(&puzzles)
	beego.Debug(n, err)
	c.Data["title"] = "Puzzle"
	puzzlez := make([]models.Puzzle, 0, len(puzzles))
	for _, it := range puzzles {
		(&it).SubString(100)
		puzzlez = append(puzzlez, it)
	}
	// fmt.Println(puzzlez)
	max, _ := models.ORM.QueryTable((*models.Puzzle)(nil)).Count()
	beego.Info(c.GetString("page"))
	c.Data["puzzles"] = puzzlez
	c.Data["pagination"] = utils.Pagination("", int(max)/15+1, page)
	c.TplName = "list.html"
}

// @router /state [get]
func (c *JudgeController) State() {
	page, _ := c.GetInt("page")
	var solutions []models.Solution
	n, err := models.ORM.QueryTable((*models.Solution)(nil)).RelatedSel().OrderBy("-created").Limit(15, 15*(page-1)).All(&solutions)
	beego.Debug(n, err)
	max, _ := models.ORM.QueryTable((*models.Solution)(nil)).Count()
	c.Data["title"] = "State"
	c.Data["solutions"] = solutions
	c.Data["pagination"] = utils.Pagination("state", int(max)/15+1, page)
	c.TplName = "state.html"
}
