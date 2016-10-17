package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/everfore/oauth/oauth2"
	"github.com/toukii/gooj/goojle/models"
	"github.com/toukii/goutils"
	"github.com/toukii/jsnm"
	"html/template"
)

var (
	OA *oauth2.OAGithub
)

func init() {
	OA = oauth2.NewOAGithub("b1499626b1250ebcea00", "261593f0356fdbdd7460929176fcb6bbe2df646a", "user", "http://goojle.daoapp.io/callback")
}

type RegistController struct {
	SessionController
}

func (c *RegistController) Prepare() {
	c.SessionController.Prepare()
}

// @router /githubsignin [get]
func (c *RegistController) Signin() {
	fmt.Println(OA.AuthURL())
	c.Redirect(OA.AuthURL(), 302)
}

// @router /callback [get]
func (c *RegistController) Callback() {
	req := c.Ctx.Request
	fmt.Printf("%s\n", req.RemoteAddr)
	b, token, err := OA.NextStepWithToken(req)
	if nil != err {
		usr := models.User{Name: "Anonymous", Passwd: "Anonymous"} //存在安全漏洞
		n := models.RegisterUser(&usr)
		if n <= 0 {
			usr := models.UserByName(usr.Name)
			n = usr.Id
		}
		c.LoginSetSession(n)
		c.Redirect("/", 302)
		return
	}
	jv := jsnm.BytesFmt(b)
	name := jv.Get("login").RawData().String()
	usr := models.User{}
	usr.Name = name
	usr.Passwd = token
	n := models.RegisterUser(&usr)
	if n <= 0 {
		usr := models.UserByName(usr.Name)
		n = usr.Id
		go func() {
			usr.Passwd = token
			_, err := models.ORM.Update(usr)
			goutils.CheckErr(err)
		}()
		c.LoginSetSession(n)
		c.Redirect("/user", 302)
		return
	}
	c.LoginSetSession(n)
	c.Redirect("/", 302)
}

// @router /register [get]
func (c *RegistController) LoadRegister() {
	c.EnableXSRF = true
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "register.html"
}

// @router /register [post]
func (c *RegistController) Register() {
	var usr models.User
	c.ParseForm(&usr)
	beego.Notice(usr)
	valid := validation.Validation{}
	usr.Valid(&valid)
	if valid.HasErrors() {
		c.Redirect("/", 302)
	}
	n := models.RegisterUser(&usr)
	if n <= 0 {
		c.Redirect("/", 302)
	}
	c.LoginSetSession(n)
	c.Redirect("/", 302)
}
