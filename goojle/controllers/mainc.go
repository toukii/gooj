package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/everfore/exc"
	"github.com/shaalx/gooj"
	"github.com/shaalx/gooj/goojle/models"
	"github.com/shaalx/goutils"
	"github.com/shaalx/jsnm"
	"html/template"
	"os"
	"strings"
	"sync"
)

type MainController struct {
	SessionController
}

var (
	defaultpath   string
	submit_LOCKER = sync.Mutex{}
	submitID      int64
)

func init() {
	defaultpath, _ = os.Getwd()
}

// @router / [get]
func (c *MainController) Get() {
	var puzzles []models.Puzzle
	n, err := models.ORM.QueryTable((*models.Puzzle)(nil)).Filter("Online", 1).All(&puzzles)
	beego.Debug(n, err)
	c.Data["title"] = "PROBLEM"
	c.Data["problems"] = puzzles
	c.TplName = "list.html"
}

// @router /problem/:id:int [get]
func (c *MainController) GetPro() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	var problem models.Puzzle
	errq := models.ORM.QueryTable((*models.Puzzle)(nil)).Filter("Id", id).RelatedSel().One(&problem)
	goutils.CheckErr(errq)
	c.Data["problem"] = problem
	addrsplit := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	rid := addrsplit[len(addrsplit)-1]
	c.Data["rid"] = rid
	c.Data["title"] = "Probs"
	c.TplName = "problem.html"
}

// @router /oj/:id:int [post]
func (c *MainController) OJ() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	fid := c.GetString("fid")
	path_ := c.GetString("rid")
	content := c.GetString("problem")
	if strings.Contains(content, `"os`) {
		c.Ctx.ResponseWriter.Write(goutils.ToByte("呵呵"))
		return
	}
	beego.Debug(content, path_, fid)

	// inser into db
	slt := models.Solution{}
	cur := c.SessionController.CurUser()
	if cur != nil {
		slt.User = &models.User{Id: cur.Id} //models.UserById(cur.Id)
	} else {
		slt.User = models.UserByName("error")
	}
	slt.Content = content
	ffid, _ := c.GetInt("fid")
	slt.Problem = models.ProblemById(ffid)
	n, dberr := models.ORM.Insert(&slt)
	if goutils.CheckErr(dberr) {

	}
	// insert into db

	submit_LOCKER.Lock()
	defer submit_LOCKER.Unlock()
	cmd := exc.NewCMD("go test -v").Cd(defaultpath)

	var puzzle models.Puzzle
	models.ORM.QueryTable((*models.Puzzle)(nil)).Filter("Id", id).One(&puzzle)
	m := Puzzle2Model(&puzzle)

	m.Content = content
	beego.Info(m)
	err := gooj.GenerateOjModle(path_, m)
	goutils.CheckErr(err)
	submitID++
	ret, err := cmd.Wd().Cd(path_).Debug().Do()
	goutils.CheckErr(err)
	result := goutils.ToString(ret)
	fmt.Println("n =", n)
	if n > 0 {
		slt.Result = result
		slt.Id = int(n)
		go func() {
			models.ORM.Update(&slt)
		}()
	}

	go cmd.Reset(fmt.Sprintf("rm -rf %s", path_)).Cd(defaultpath).Execute()

	c.Ctx.ResponseWriter.Write(ret)
}

func Puzzle2Model(puzzle *models.Puzzle) *gooj.Model {
	if nil == puzzle {
		return nil
	}
	m := gooj.Model{}
	m.Id = fmt.Sprintf("%s", puzzle.Id)
	m.ArgsType = puzzle.ArgsType
	m.Content = puzzle.Content
	m.Desc = puzzle.Descr
	m.FuncName = puzzle.FuncName
	m.Online = puzzle.Online
	m.RetsType = puzzle.RetsType
	m.TestCases = puzzle.TestCases
	m.Title = puzzle.Title
	return &m
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
		c.Redirect("/", 302)
	}
	c.LoginSetSession(uid)
	c.Redirect("/user", 302)
}

func (c *MainController) Prepare() {
	c.SessionController.Prepare()
}

func (c *MainController) LoginSetSession(usrid int) {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
		c.Redirect("/", 302)
	}

	sess.Set("gosessionid", usrid)
	beego.Debug("set [gosessionid]----->", usrid)
	sess.SessionRelease(c.Ctx.ResponseWriter)
}

func (c *MainController) LogoutSetSession() {
	sess, err := models.GlobalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil || sess == nil {
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
	c.Redirect("/", 302)
}

// @router /user [get]
func (c *MainController) User() {
	cur := c.SessionController.CurUser()
	if cur != nil {
		var probs []models.Problem
		models.ORM.QueryTable((*models.Problem)(nil)).Filter("User__Id", cur.Id).Limit(5).All(&probs)
		problemz := make([]gooj.Model, 0, len(probs))
		for _, it := range probs {
			js := jsnm.BytesFmt(goutils.ToByte(it.Content))
			md := gooj.Model{}
			md.Id = js.Get("id").RawData().String()
			md.Desc = js.Get("desc").RawData().String()
			md.Title = js.Get("title").RawData().String()
			problemz = append(problemz, md)
		}
		c.Data["problems"] = problemz

		var solutions []models.Solution
		models.ORM.QueryTable((*models.Solution)(nil)).Filter("User__Id", cur.Id).Limit(5).OrderBy("-Id").All(&solutions)
		c.Data["solutions"] = solutions
		fmt.Println(solutions)
	}

	c.TplName = "user.html"
}
