package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterDataBase("localhost", "mysql", "root:1234@tcp(localhost:3306)/session?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "uEIYt69vVZXOJDok:pQCautZUg6ji0fdmL@tcp(10.10.26.58:3306)/k3MHIXiO61hr5vC0?charset=utf8")
	// orm.RegisterDataBase("default", "mysql", "goojle:Goojle123@tcp(121.42.161.248:3306)/gooj?charset=utf8")
	orm.RegisterModel(new(User), new(Puzzle), new(Solution), new(Remark), new(Result))
	ORM = orm.NewOrm()
	ORM.Using("default")
	orm.Debug = true
}

var ORM orm.Ormer

func RegisterUser(usr *User) int {
	n, err := ORM.Insert(usr)
	if err != nil {
		return -1
	}
	return int(n)
}

func CheckUser(usr *User) int {
	err := ORM.QueryTable(usr).Filter("Name", usr.Name).Filter("Passwd", usr.Passwd).One(usr)
	if err != nil {
		return -1
	}
	return usr.Id
}

func RemarksById(id int) []Remark {
	var remarks []Remark
	_, err := ORM.QueryTable((*Remark)(nil)).Filter("Topic__Id", id).RelatedSel("User").All(&remarks)
	if err != nil {
		return nil
	}
	return remarks
}

func UserByName(name string) *User {
	var usr User
	if err := ORM.QueryTable((*User)(nil)).Filter("Name", name).One(&usr); err != nil {
		fmt.Println(err, name)
		return nil
	}
	if usr.Id < 0 {
		return nil
	}
	fmt.Println(usr)
	return &usr
}

func UserById(id int) *User {
	var usr User
	if err := ORM.QueryTable((*User)(nil)).Filter("Id", id).One(&usr); err != nil {
		fmt.Println(err, id)
		return nil
	}
	if usr.Id < 0 {
		return nil
	}
	fmt.Println(usr)
	return &usr
}

func DelRemardById(id int) bool {
	n, err := ORM.QueryTable((*Remark)(nil)).Filter("Id", id).Delete()
	if n <= 0 || err != nil {
		return false
	}
	return true
}

func RemarksByUserId(userid int) []Remark {
	var remarks []Remark
	// if _, err := ORM.QueryTable((*Remark)(nil)).Filter("User__Id", userid).RelatedSel("topic").All(&remarks); err != nil {
	if _, err := ORM.QueryTable((*Remark)(nil)).Filter("User__Id", userid).RelatedSel().All(&remarks); err != nil {
		fmt.Println(err)
		return nil
	}
	return remarks
}

func RemarkById(remarkid int) *Remark {
	var remark Remark
	if err := ORM.QueryTable((*Remark)(nil)).Filter("Id", remarkid).RelatedSel().One(&remark); err != nil {
		return nil
	}
	return &remark
}
