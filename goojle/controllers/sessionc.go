package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/shaalx/gooj/goojle/models"
	"strconv"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) Prepare() {
	c.SetCurUser()
}

func (c *SessionController) LoginSetSession(usrid int) {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		// c.Abort("401")
		c.Redirect("/", 302)
	}

	sess.Set("gosessionid", usrid)
	beego.Debug("set [gosessionid]----->", usrid)
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

func (c *SessionController) LogoutSetSession() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		// c.Abort("401")
		c.Redirect("/", 302)
	}
	sess.Set("gosessionid", "_")
	beego.Debug("set [gosessionid]-----> _")
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

func (c *SessionController) CurUser() *models.User {
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

func (c *SessionController) SetCurUser() {
	user := c.CurUser()
	if user != nil {
		c.Data["curUser"] = user
	}
}
