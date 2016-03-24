package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/shaalx/gooj/goojle/models"
	"html/template"
	"strconv"
	"strings"
)

type LogController struct {
	beego.Controller
}

// @router /login [get]
func (c *LogController) LoadLogin() {
	c.EnableXSRF = true
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login.html"
}

// @router /login [post]
func (c *LogController) Login() {
	var usr models.User
	err := c.ParseForm(&usr)
	beego.Debug("login user:", usr, err)
	c.Data["curUser"] = &usr
	if err != nil {
		c.Abort("403")
	}
	uid := usr.Check()
	if uid <= 0 {
		c.Abort("401")
	}
	c.LoginSetSession(uid)
	c.Get()
}

func (c *LogController) Prepare() {
	user := c.CurUser()
	c.Data["curUser"] = user
	uri := c.Ctx.Request.RequestURI
	if strings.EqualFold(uri, "/publish") || strings.Contains(uri, "/remark") || strings.Contains(uri, "/del") {
		c.CheckLogin()
	}
}

func (c *LogController) CheckLogin() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Abort("401")
	}
	sessid := sess.Get("gosessionid")
	beego.Debug(sessid)
	if sessid == nil || strings.Contains(fmt.Sprintf("%v", sessid), "_") {
		c.Abort("401")
	}
}

func (c *LogController) LoginSetSession(usrid int) {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Abort("401")
	}

	sess.Set("gosessionid", usrid)
	beego.Debug("set [gosessionid]----->", usrid)
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

func (c *LogController) LogoutSetSession() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Abort("401")
	}
	sess.Set("gosessionid", "_")
	beego.Debug("set [gosessionid]-----> _")
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

// @router /logout [get]
func (c *LogController) Logout() {
	c.Data["curUser"] = nil
	c.LogoutSetSession()
	c.Get()
}

func (c *LogController) CurUser() *models.User {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		return nil
	}

	iuserid := sess.Get("gosessionid")
	beego.Debug("get [gosessionid] <------- ", iuserid)
	if iuserid == nil {
		return nil
	}
	userid := fmt.Sprintf("%v", iuserid)
	id, err := strconv.Atoi(userid)
	if err != nil {
		return nil
	}
	if id <= 0 {
		return nil
	}
	usr := models.UserById(id)
	if nil == usr {
		return nil
	}
	beego.Debug("current user ----> ", *usr)
	return usr
}

// @router /user [get]
func (c *LogController) User() {
	user := c.CurUser()
	if user == nil {
		c.Abort("401")
	}
	c.Data["user"] = user
	c.Data["topics"] = models.TopicsById(user.Id)
	c.Data["remarks"] = models.RemarksByUserId(user.Id)
	c.TplName = "user.html"
}
