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

var (
	tpl         map[string]string
	defaultpath string
	m           *gooj.Model
	proMap      map[string]gooj.Model
	pro1        = `package goojt

func reverse(arg []int) []int {
	// TODO Something
	return nil
}`
	scripts = `<link href="http://cdn.bootcss.com/bootstrap/3.3.4/css/bootstrap.min.css" rel="stylesheet">
    <link href="http://cdn.bootcss.com/font-awesome/4.2.0/css/font-awesome.min.css" rel="stylesheet">
    <link href="http://static.bootcss.com/www/assets/css/site.min.css?v5" rel="stylesheet">
    <script src="http://cdn.bootcss.com/jquery/1.11.2/jquery.min.js"></script>
    <script src="http://cdn.bootcss.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
    <script src="http://cdn.bootcss.com/unveil/1.3.0/jquery.unveil.min.js"></script>
    <script src="http://cdn.bootcss.com/scrollup/2.4.0/jquery.scrollUp.min.js"></script>
    <script src="http://cdn.bootcss.com/toc/0.3.2/toc.min.js"></script>
    <script src="http://cdn.bootcss.com/jquery.matchHeight/0.5.2/jquery.matchHeight-min.js"></script>
    <script src="http://static.bootcss.com/www/assets/js/site.min.js"></script>`
	script template.HTML
	msURL  = "http://7xku3c.com1.z0.glb.clouddn.com/models.json"
)

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
	script = template.HTML(scripts)
	ms := gooj.TiniuMs(msURL)
	proMap = make(map[string]gooj.Model)
	for _, it := range ms {
		proMap[it.Id] = it
	}
}

func main() {
	http.HandleFunc("/pro", pro)
	http.HandleFunc("/", list)
	http.HandleFunc("/oj", submit)
	http.ListenAndServe(":80", nil)
}

func list(rw http.ResponseWriter, req *http.Request) {
	tpl, err := template.New("list.html").Parse(tpl["list"])
	goutils.CheckErr(err)
	data := make(map[string]interface{})
	data["pros"] = gooj.TiniuMs(msURL)
	data["script"] = script
	tpl.Execute(rw, data)
}

func pro(rw http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI
	uris := strings.Split(uri, "/")
	length := len(uris)
	log.Info(uris[len(uris)-1])
	log.Info(uri)
	if length <= 0 {
		return
	}
	m := proMap[uris[len(uris)-1]]
	tpl, err := template.New("pro.html").Parse(tpl["pro"])
	goutils.CheckErr(err)
	data := make(map[string]interface{})
	data["script"] = script
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
	m.Content = content
	err := gooj.GenerateOjModle(path_, m)
	goutils.CheckErr(err)
	ret, err := cmd.Cd(path_).Debug().Do()
	goutils.CheckErr(err)
	rw.Write(ret)
}

func rand() {

}
