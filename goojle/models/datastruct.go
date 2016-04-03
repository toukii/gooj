package models

import (
	"github.com/astaxie/beego/validation"
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
	v.Required(u.Name, "shaalx")
	if u.Passwd != u.RePasswd {
		v.SetError("passwd", "repassword does not equal.")
	}
}

func (u *User) Check() int {
	return CheckUser(u)
}

type Topic struct {
	Id      int       `orm:"id;pk" form:"-"`
	User    *User     `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Title   string    `orm:"title" form:"title"`
	Content string    `orm:"content;null" form:"content"`
	Create  time.Time `orm:"auto_now_add;column(created);type(datetime)"`
}

func (t *Topic) Publish() error {
	return PublishTopic(t)
}

/*create table remark(
	id int auto_increment primary key,
	user_id int not null,
	solution_id int not null,
	content text,
	created datetime
)*/
type Remark struct {
	Id      int       `orm:"id;pk" form:"-"`
	User    *User     `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Problem *Problem  `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Content string    `orm:"content" form:"content"`
	Topic   *Topic    `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Create  time.Time `orm:"auto_now_add;column(created);type(datetime)"`
}

func (r *Remark) Publish() error {
	return PublishRemark(r)
}

type Problem struct {
	Id      int    `orm:"id;pk"`
	User    *User  `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Content string `orm:"content" form:"content"`
}

/*create table solution(
	id int auto_increment primary key,
	user_id int not null,
	problem_id char(100) not null,
	content text,
	created datetime
)*/
type Solution struct {
	Id      int       `orm:"id;pk"`
	User    *User     `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Problem *Problem  `orm:"rel(fk);null;on_delete(cascade)" form:"-"`
	Content string    `orm:"content;null" form:"content"`
	Result  string    `orm:"result;null" form:"result"`
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
