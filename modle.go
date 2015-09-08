package gooj

import (
	"fmt"
	"github.com/shaalx/goutils"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

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
	return renderOjModle(args, test_file)
}

func generateOjFunc(content string, w io.Writer) error {
	_, err := w.Write(goutils.ToByte(content))
	return err
}
