package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/everfore/oauth/oauth2"
	"github.com/shaalx/gooj/goojle/models"
	"github.com/shaalx/jsnm"
	"html/template"
)

var (
	OA *oauth2.OAGithub
)

func init() {
	OA = oauth2.NewOAGithub("8dbf243923c1384ebb63", "f7129496415cefed1770e0d2470d14cd82015f25", "user", "http://goojle.daoapp.io/callback")
}

type LogController struct {
	// beego.Controller
	SessionController
}

func (c *LogController) Prepare() {
	c.SessionController.Prepare()
}

// @router /githubsignin [get]
func (c *LogController) Signin() {
	fmt.Println(OA.AuthURL())
	c.Redirect(OA.AuthURL(), 302)
}

// @router /callback [get]
func (c *LogController) Callback() {
	req := c.Ctx.Request
	// rw := c.Ctx.ResponseWriter
	fmt.Printf("%s\n", req.RemoteAddr)
	b, err := OA.NextStep(req)
	if nil != err {
		usr := models.User{Name: "error", Passwd: "error"}
		n := models.RegisterUser(&usr)
		if n <= 0 {
			usr := models.UserByName(usr.Name)
			n = usr.Id
		}
		c.LoginSetSession(n)
		c.Redirect("/", 302)
		// rw.Write([]byte(err.Error()))
		return
	}
	jv := jsnm.BytesFmt(b)
	name := jv.Get("login").RawData().String()
	usr := models.User{}
	usr.Name = name
	usr.Passwd = name
	n := models.RegisterUser(&usr)
	if n <= 0 {
		usr := models.UserByName(usr.Name)
		n = usr.Id
	}
	c.LoginSetSession(n)
	fmt.Print(jv.MapData())
	c.Redirect("/", 302)
}

// @router /register [get]
func (c *LogController) LoadRegister() {
	c.EnableXSRF = true
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "register.html"
}

// @router /register [post]
func (c *LogController) Register() {
	var usr models.User
	c.ParseForm(&usr)
	beego.Notice(usr)
	valid := validation.Validation{}
	usr.Valid(&valid)
	if valid.HasErrors() {
		// c.Abort("401")
		c.Redirect("/", 302)
	}
	n := models.RegisterUser(&usr)
	if n <= 0 {
		// c.Abort("401")
		c.Redirect("/", 302)
	}
	c.LoginSetSession(n)
	c.Redirect("/", 302)
}
