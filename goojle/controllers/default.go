package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/everfore/exc"
	"github.com/shaalx/gooj"
	"github.com/shaalx/goutils"
	"net/http"
	"os"
	"strings"
)

type MainController struct {
	beego.Controller
}

var (
	problemURL  = "http://7xku3c.com1.z0.glb.clouddn.com/models.json"
	problems    []gooj.Model
	problemMap  map[string]gooj.Model
	defaultpath string
)

func init() {
	defaultpath, _ = os.Getwd()
	problems = gooj.TiniuMs(problemURL)
	problemMap = make(map[string]gooj.Model)
	for _, it := range problems {
		problemMap[it.Id] = it
	}
}

// @router /update [get]
func (c *MainController) Update() {
	problems = gooj.TiniuMs(problemURL)
	problemMap = make(map[string]gooj.Model)
	for _, it := range problems {
		problemMap[it.Id] = it
	}
	c.Redirect("/", 302)
}

// @router / [get]
func (c *MainController) Get() {
	c.Data["title"] = "PROBLEM"
	c.Data["problems"] = problems
	c.TplNames = "list.html"
}

// @router /problem/:id:int [get]
func (c *MainController) GetPro() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	c.Data["problem"] = problemMap[fmt.Sprintf("%d", id)]
	c.TplNames = "problem.html"
}

type Input struct {
	fname   string `form:"fname"`
	problem string `form:"problem"`
	rid     string `form:"rid"`
}

// @router /oj/:id:int [post]
func (c *MainController) OJ() {
	var input Input
	c.ParseForm(&input)
	beego.Info(input)
	beego.Info(c.Ctx.Input.Param("problem"))
	submit(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func submit(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fid := req.Form.Get("fid")
	path_ := req.Form.Get("rid")
	content := req.Form.Get("problem")
	if strings.Contains(content, `"os`) {
		rw.Write(goutils.ToByte("呵呵"))
		return
	}
	beego.Debug(content, path_, fid)
	cmd := exc.NewCMD("go test -v").Cd(defaultpath)
	m := problemMap[fid]
	m.Content = content
	beego.Info(m)
	err := gooj.GenerateOjModle(path_, &m)
	goutils.CheckErr(err)
	ret, err := cmd.Wd().Cd(path_).Debug().Do()
	goutils.CheckErr(err)
	rw.Write(ret)
}
