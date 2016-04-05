package model_util

import (
	"fmt"
	"github.com/shaalx/goutils"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type Model struct {
	Id        string `from:"id" json:"id"`
	Title     string `from:"title" json:"title"`
	Desc      string `from:"descr" json:"desc"`
	FuncName  string `from:"func_name" json:"func_name"`
	Content   string `from:"content" json:"content"`
	ArgsType  string `from:"args_type" json:"args_type"`
	RetsType  string `from:"rets_type" json:"rets_type"`
	TestCases string `from:"test_cases" json:"test_cases"`
	Online    byte   `from:"online" json:"online"`
}

func renderOjModle(args map[string]interface{}, w io.Writer) error {
	b, err := ioutil.ReadFile("oj.tpl")
	if goutils.CheckErr(err) {
		return err
	}
	s := goutils.ToString(b)
	tpl, err := template.New("oj.tpl").Parse(s)
	if goutils.CheckErr(err) {
		return err
	}
	return tpl.Execute(w, args)
}

func GenerateOjModle(path_ string, m *Model) error {
	_, err := os.Stat(path_)
	if !goutils.CheckErr(err) {
		os.RemoveAll(path_)
	}
	err = os.Mkdir(path_, 0777)
	if goutils.CheckErr(err) {
		// return err
	}
	fname := m.FuncName + ".go"
	tname := m.FuncName + "_test.go"
	func_file, err := os.OpenFile(filepath.Join(path_, fname), os.O_CREATE|os.O_WRONLY, 0644)
	defer func_file.Close()
	if goutils.CheckErr(err) {
		return err
	}
	if goutils.CheckErr(generateOjFunc(m.Content, func_file)) {
		return fmt.Errorf("generateOjFunc error")
	}

	test_file, err := os.OpenFile(filepath.Join(path_, tname), os.O_CREATE|os.O_WRONLY, 0644)
	defer test_file.Close()
	if goutils.CheckErr(err) {
		return err
	}
	args := make(map[string]interface{})
	args["FUNC"] = m.FuncName
	args["ArgsType"] = m.ArgsType
	args["RetsType"] = m.RetsType
	args["TestCases"] = m.TestCases
	return renderOjModle(args, test_file)
}

func generateOjFunc(content string, w io.Writer) error {
	_, err := w.Write(goutils.ToByte(content))
	return err
}
