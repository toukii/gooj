package controllers

import (
	"fmt"
	"github.com/shaalx/gooj/goojle/models"
	// "github.com/shaalx/goutils"
)

type UserController struct {
	SessionController
}

func (c *UserController) Prepare() {
	c.SessionController.Prepare()
}

// @router /user [get]
func (c *UserController) User() {
	cur := c.CurUser()
	if cur != nil {
		puzzles_count, _ := models.ORM.QueryTable((*models.Puzzle)(nil)).Filter("User__Id", cur.Id).Count()
		c.Data["puzzles_count"] = puzzles_count

		solutions_count, _ := models.ORM.QueryTable((*models.Solution)(nil)).Filter("User__Id", cur.Id).Count()
		c.Data["solutions_count"] = solutions_count
	}

	c.TplName = "usr/user.html"
}

// @router /puzzles [get]
func (c *UserController) Puzzles() {
	cur := c.CurUser()
	if cur != nil {
		var puzzles []models.Puzzle
		models.ORM.QueryTable((*models.Puzzle)(nil)).Filter("User__Id", cur.Id).RelatedSel().Limit(15).OrderBy("-Id").All(&puzzles)
		c.Data["puzzles"] = puzzles
	}

	c.TplName = "usr/puzzles.html"
}

// @router /solutions [get]
func (c *UserController) Solutions() {
	cur := c.CurUser()
	if cur != nil {
		var solutions []models.Solution
		models.ORM.QueryTable((*models.Solution)(nil)).Filter("User__Id", cur.Id).Limit(15).RelatedSel().OrderBy("-Id").All(&solutions)
		c.Data["solutions"] = solutions
		fmt.Println(solutions)
	}

	c.TplName = "usr/solutions.html"
}
