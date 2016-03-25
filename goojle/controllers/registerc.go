package controllers

// import (
// 	"github.com/astaxie/beego"
// 	"github.com/astaxie/beego/validation"
// 	"github.com/shaalx/gooj/goojle/models"
// 	"html/template"
// )

// type RegisterController struct {
// 	LogController
// }

// // @router /register [get]
// func (c *RegisterController) LoadRegister() {
// 	c.EnableXSRF = true
// 	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
// 	c.TplName = "register.html"
// }

// // @router /register [post]
// func (c *RegisterController) Register() {
// 	var usr models.User
// 	c.ParseForm(&usr)
// 	beego.Notice(usr)
// 	valid := validation.Validation{}
// 	usr.Valid(&valid)
// 	if valid.HasErrors() {
// 		c.Abort("401")
// 	}
// 	n := models.RegisterUser(&usr)
// 	if n <= 0 {
// 		c.Abort("401")
// 	}
// 	c.LogController.LoginSetSession(n)
// 	c.LogController.Get()
// }
