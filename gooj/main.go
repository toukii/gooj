package main

import (
	"github.com/everfore/exc"
	"github.com/qiniu/log"
	"github.com/shaalx/gooj"
	"github.com/shaalx/goutils"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var tpl map[string]string
var defaultpath string
var m *gooj.Model
var ms []gooj.Model
var pro1 = `package goojt

func reverse(arg []int) []int {
	// TODO Something
	return nil
}`

func init() {
	defaultpath, _ = os.Getwd()
	tpl = make(map[string]string)
	log.SetOutputLevel(log.Ldebug)
	b, err := ioutil.ReadFile("pro.html")
	goutils.CheckErr(err)
	tpl["pro"] = goutils.ToString(b)
	b, err = ioutil.ReadFile("list.html")
	goutils.CheckErr(err)
	tpl["list"] = goutils.ToString(b)
	m = gooj.ToM()
	ms = gooj.ToMs()
}

func main() {
	http.HandleFunc("/", pro)
	http.HandleFunc("/l", list)
	http.HandleFunc("/oj", submit)
	http.ListenAndServe(":80", nil)
}

func list(rw http.ResponseWriter, req *http.Request) {
	tpl, err := template.New("list.html").Parse(tpl["list"])
	goutils.CheckErr(err)
	data := make(map[string]interface{})
	data["pros"] = ms
	tpl.Execute(rw, data)
}

func pro(rw http.ResponseWriter, req *http.Request) {
	tpl, err := template.New("pro.html").Parse(tpl["pro"])
	goutils.CheckErr(err)
	data := make(map[string]interface{})
	data["pro"] = m.Content
	data["desc"] = m.Desc
	data["fname"] = m.FuncName
	tpl.Execute(rw, data)
}

func submit(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	path_ := req.Form.Get("rid")
	// func_name := req.Form.Get("fname")
	content := req.Form.Get("pro")
	if strings.Contains(content, `"os`) {
		rw.Write(goutils.ToByte("呵呵"))
		return
	}
	cmd := exc.NewCMD("go test -v").Cd(defaultpath)
	err := gooj.GenerateOjModle(path_, m)
	goutils.CheckErr(err)
	ret, err := cmd.Cd(path_).Debug().Do()
	goutils.CheckErr(err)
	rw.Write(ret)
}

func rand() {

}
