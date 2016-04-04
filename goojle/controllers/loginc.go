package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/shaalx/gooj/goojle/models"
	// "github.com/shaalx/goutils"
	"html/template"
)

type LoginController struct {
	SessionController
}

func (c *LoginController) Prepare() {
	c.SessionController.Prepare()
}

// @router /login [get]
func (c *LoginController) LoadLogin() {
	c.EnableXSRF = true
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login.html"
}

// @router /login [post]
func (c *LoginController) Login() {
	var usr models.User
	err := c.ParseForm(&usr)
	beego.Debug("login user:", usr, err)
	c.Data["curUser"] = &usr
	if err != nil {
		c.Abort("403")
	}
	uid := usr.Check()
	if uid <= 0 {
		c.Redirect("/", 302)
	}
	c.LoginSetSession(uid)
	c.Redirect("/user", 302)
}

func (c *LoginController) LoginSetSession(usrid int) {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Redirect("/", 302)
	}

	sess.Set("gosessionid", usrid)
	beego.Debug("set [gosessionid]----->", usrid)
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

func (c *LoginController) LogoutSetSession() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Redirect("/", 302)
	}
	sess.Set("gosessionid", "_")
	beego.Debug("set [gosessionid]-----> _")
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

// @router /logout [get]
func (c *LoginController) Logout() {
	c.Data["curUser"] = nil
	c.LogoutSetSession()
	c.Redirect("/", 302)
}

// @router /user [get]
func (c *LoginController) User() {
	cur := c.CurUser()
	if cur != nil {
		var puzzles []models.Puzzle
		models.ORM.QueryTable((*models.Puzzle)(nil)).Filter("User__Id", cur.Id).Limit(5).All(&puzzles)
		c.Data["puzzles"] = puzzles

		var solutions []models.Solution
		models.ORM.QueryTable((*models.Solution)(nil)).Filter("User__Id", cur.Id).Limit(5).OrderBy("-Id").All(&solutions)
		c.Data["solutions"] = solutions
		fmt.Println(solutions)
	}

	c.TplName = "user.html"
}
