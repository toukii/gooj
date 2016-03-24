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

	"sync"
)

type MainController struct {
	beego.Controller
}

var (
	// problemURL    = "http://7xku3c.com1.z0.glb.clouddn.com/models.json"
	problemURL    = "https://raw.githubusercontent.com/shaalx/GoOJProbs/master/models.json"
	problems      []gooj.Model
	problemMap    map[string]gooj.Model
	defaultpath   string
	submit_LOCKER = sync.Mutex{}
	submitID      int64
)

func init() {
	defaultpath, _ = os.Getwd()
	// problems = gooj.TiniuMs(problemURL)
	problems = gooj.ToMs()
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
	c.TplName = "list.html"
}

// @router /problem/:id:int [get]
func (c *MainController) GetPro() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	c.Data["problem"] = problemMap[fmt.Sprintf("%d", id)]
	addrsplit := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	rid := addrsplit[len(addrsplit)-1]
	c.Data["rid"] = rid
	c.Data["title"] = "Probs"
	c.TplName = "problem.html"
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
	beego.Info("problem******************:", c.Ctx.Input.Param("problem"))
	res := submit(c.Ctx.ResponseWriter, c.Ctx.Request)
	c.TplName = "result.html"
	c.Data["result"] = goutils.ToString(res)
}

func submit(rw http.ResponseWriter, req *http.Request) []byte {
	req.ParseForm()
	fid := req.Form.Get("fid")
	path_ := req.Form.Get("rid")
	content := req.Form.Get("problem")
	if strings.Contains(content, `"os`) {
		rw.Write(goutils.ToByte("呵呵"))
		return goutils.ToByte("呵呵")
	}
	beego.Debug(content, path_, fid)
	submit_LOCKER.Lock()
	defer submit_LOCKER.Unlock()
	cmd := exc.NewCMD("go test -v").Cd(defaultpath)
	m := problemMap[fid]
	m.Content = content
	beego.Info(m)
	err := gooj.GenerateOjModle(path_, &m)
	goutils.CheckErr(err)
	submitID++
	ret, err := cmd.Wd().Cd(path_).Debug().Do()
	goutils.CheckErr(err)
	rw.Write(ret)
	fmt.Println(goutils.ToString(ret))
	go cmd.Reset(fmt.Sprintf("rm -rf %s", path_)).Cd(defaultpath).ExecuteAfter(5)
	return ret
}
