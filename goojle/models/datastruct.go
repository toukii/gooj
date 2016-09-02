package models

import (
	"github.com/astaxie/beego/validation"
	"github.com/toukii/gooj/goojle/utils"
	"time"
)

type User struct {
	Id       int    `orm:"id;pk" form:"-"`
	Name     string `form:"username" orm:"unique"`
	Passwd   string `form:"password" orm:"passwd"`
	RePasswd string `form:"repassword" orm:"-"`
	Xsrf     string `form:"_xsrf" orm:"-"`
}

func (u *User) Valid(v *validation.Validation) {
	v.Required(u.Name, "toukii")
	if u.Passwd != u.RePasswd {
		v.SetError("passwd", "repassword does not equal.")
	}
}

func (u *User) Check() int {
	return CheckUser(u)
}

type Remark struct {
	Id       int       `orm:"id;pk" form:"-"`
	User     *User     `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Puzzle   *Puzzle   `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Solution *Solution `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Content  string    `orm:"content" form:"content"`
	Create   time.Time `orm:"auto_now_add;column(created);type(datetime)"`
}

type Solution struct {
	Id      int       `orm:"id;pk"`
	User    *User     `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Puzzle  *Puzzle   `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Content string    `orm:"content;null" form:"content"`
	Result  *Result   `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Create  time.Time `orm:"auto_now_add;column(created);type(datetime)"`
}

type Puzzle struct {
	Id        int    `orm:"id;pk" form:"id"`
	User      *User  `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Title     string `json:"title" orm:"title" form:"title"`
	Descr     string `json:"descr" orm:"descr" form:"descr"`
	FuncName  string `json:"func_name" orm:"func_name" form:"func_name"`
	Content   string `json:"content" orm:"content" form:"content"`
	ArgsType  string `json:"args_type" orm:"args_type" form:"args_type"`
	RetsType  string `json:"rets_type" orm:"rets_type" form:"rets_type"`
	TestCases string `json:"test_cases" orm:"test_cases" form:"test_cases"`
	Online    byte   `json:"online" orm:"online" form:"online"`
}

func (p *Puzzle) SubString(length int) {
	p.Descr = utils.SubString(p.Descr, length)
}

type Result struct {
	Id          int    `orm:"id;pk" form:"id"`
	State       string `json:"state" orm:"state" form:"state"`
	RunCostTime string `json:"run_cost_time" orm:"run_cost_time" form:"run_cost_time"`
	TestCase    string `json:"test_case" orm:"test_case" form:"test_case"`
	RunResult   string `json:"run_result" orm:"run_result" form:"run_result"`
	ErrorInfo   string `json:"error_info" orm:"error_info" form:"error_info"`
	Content     string `json:"content" orm:"content" form:"content"`
}

func AnalyseResultParse(res *utils.Result) *Result {
	ret := Result{}
	ret.State = res.State
	ret.RunCostTime = res.RunCostTime
	ret.TestCase = res.Fail.TestCase
	ret.RunResult = res.Fail.RunResult
	ret.ErrorInfo = res.Fail.ErrorInfo
	ret.Content = res.Content
	return &ret
}
