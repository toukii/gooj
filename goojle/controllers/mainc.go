package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/everfore/exc"
	"github.com/shaalx/gooj"
	"github.com/shaalx/gooj/goojle/models"
	"github.com/shaalx/goutils"
	"html/template"
	"net/http"
	"os"
	"strings"
	"sync"
)

type MainController struct {
	// beego.Controller
	SessionController
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

// @router /login [get]
func (c *MainController) LoadLogin() {
	c.EnableXSRF = true
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login.html"
}

// @router /login [post]
func (c *MainController) Login() {
	var usr models.User
	err := c.ParseForm(&usr)
	beego.Debug("login user:", usr, err)
	c.Data["curUser"] = &usr
	if err != nil {
		c.Abort("403")
	}
	uid := usr.Check()
	if uid <= 0 {
		// c.Abort("401")
		c.Redirect("/", 302)
	}
	c.LoginSetSession(uid)
	c.Data["title"] = "PROBLEM"
	c.Data["problems"] = problems
	c.TplName = "list.html"
	// c.Get()
	// c.Redirect("/", 302)
}

func (c *MainController) Prepare() {
	c.SessionController.Prepare()
	c.CheckLogin()
}

func (c *MainController) CheckLogin() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		beego.Debug(err)
	}
	sessid := sess.Get("gosessionid")
	beego.Debug(sessid)
}

func (c *MainController) LoginSetSession(usrid int) {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		// c.Abort("401")
		c.Redirect("/", 302)
	}

	sess.Set("gosessionid", usrid)
	beego.Debug("set [gosessionid]----->", usrid)
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

func (c *MainController) LogoutSetSession() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		// c.Abort("401")
		c.Redirect("/", 302)
	}
	sess.Set("gosessionid", "_")
	beego.Debug("set [gosessionid]-----> _")
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

// @router /logout [get]
func (c *MainController) Logout() {
	c.Data["curUser"] = nil
	c.LogoutSetSession()
	// c.Get()
	c.Redirect("/", 302)
}

// @router /user [get]
func (c *MainController) User() {
	c.TplName = "user.html"
}
