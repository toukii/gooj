package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/everfore/exc"
	"github.com/shaalx/gooj/goojle/models"
	"github.com/shaalx/gooj/goojle/utils"
	"github.com/shaalx/gooj/model_util"
	"github.com/shaalx/goutils"
	"html/template"
	"os"
	"strings"
	"sync"
)

type ListController struct {
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

func ojCheck(id int) int {
	n, err := models.ORM.QueryTable((*models.Puzzle)(nil)).Filter("Id", id).Filter("Online", 1).Count()
	if goutils.CheckErr(err) {
		return 0
	}
	return int(n)
}

// @router /oj/:id:int [get]
func (c *ListController) GetPro() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	var puzzle models.Puzzle
	puzzle.Id = id // 为了在页面上渲染
	if ojCheck(id) <= 0 {
		c.Data["Content"] = template.HTML("<h3>题目还未审核，如有问题，请联系管理员。</h3>")
		c.Data["puzzle"] = puzzle
		c.TplName = "sorry.html"
		return
	}
	errq := models.ORM.QueryTable((*models.Puzzle)(nil)).Filter("Id", id).RelatedSel().One(&puzzle)
	goutils.CheckErr(errq)
	c.Data["puzzle"] = puzzle
	addrsplit := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	rid := addrsplit[len(addrsplit)-1]
	c.Data["rid"] = rid
	c.Data["title"] = "Probs"
	c.TplName = "oj.html"
}

// @router /oj/:id:int [post]
func (c *ListController) OJ() {
	var id int
	c.Ctx.Input.Bind(&id, ":id")
	if ojCheck(id) <= 0 {
		c.Data["Content"] = template.HTML("<h3>题目还未审核，如有问题，请联系管理员。</h3>")
		c.TplName = "sorry.html"
		return
	}
	fid := c.GetString("fid")
	path_ := c.GetString("rid")
	content := c.GetString("puzzle")
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
		slt.User = models.UserByName("Anonymous")
	}
	slt.Content = content
	ffid, _ := c.GetInt("fid")
	slt.Puzzle = &models.Puzzle{Id: ffid}
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
	err := model_util.GenerateOjModle(path_, m)
	goutils.CheckErr(err)
	submitID++
	test_result, err := cmd.Wd().Cd(path_).Debug().Do()
	analyse_result := utils.Analyse(goutils.ToString(test_result))
	fmt.Println(analyse_result)
	goutils.CheckErr(err)
	if n > 0 {
		result := models.AnalyseResultParse(analyse_result)
		n, err := models.ORM.Insert(result)
		if !goutils.CheckErr(err) {
			result.Id = int(n)
		}
		slt.Result = result
		slt.Id = int(n)
		go func() {
			models.ORM.Update(&slt)
		}()
	}

	go cmd.Reset(fmt.Sprintf("rm -rf %s", path_)).Cd(defaultpath).Execute()

	c.Ctx.ResponseWriter.Write(goutils.ToByte(analyse_result.String()))
}

func Puzzle2Model(puzzle *models.Puzzle) *model_util.Model {
	if nil == puzzle {
		return nil
	}
	m := model_util.Model{}
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
