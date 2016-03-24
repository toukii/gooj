package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Errorlogin() {
	c.EnableXSRF = true
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login.html"
}
