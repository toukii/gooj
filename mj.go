package gooj

import (
	"encoding/json"
	"github.com/shaalx/goutils"
	"io/ioutil"
	"net/http"
	"os"
)

/*
{
"desc":"reverse the array",
"func_model":"package goojt"+"\n"+
"func reverse(arr []int) []int{"+"\n"+
"return nil"+"\n"+
"}",

"func_name":"reverse",
"args":"[]int",
"rets":"[]int",
"test_cases":"{in: []int{1, 2, 3}, out: []int{3, 2, 1}},"+"\n"+
		"{in: []int{1, 2, 4}, out: []int{4, 2, 1}},"+"\n"+
		"{in: []int{1, 5, 3}, out: []int{3, 5, 1}},"+"\n"+
		"{in: []int{6, 2, 3}, out: []int{3, 2, 6}},",
}*/

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

func NewModel() Model {
	return Model{
		Id:       "1",
		Title:    "reverse",
		Desc:     "reverse the array",
		FuncName: "reverse",
		Content: `package goojt

func reverse(in []int) []int {
	leng := len(in)
	l := leng / 2
	for i := 0; i < l; i++ {
		in[i], in[leng-1-i] = in[leng-1-i], in[i]
	}
	return in
}`,
		ArgsType: "[]int",
		RetsType: "[]int",
		TestCases: `{in: []int{1, 2, 3}, out: []int{3, 2, 1}},
		{in: []int{1, 2, 4}, out: []int{4, 2, 1}},
		{in: []int{1, 5, 3}, out: []int{3, 5, 1}},
		{in: []int{6, 2, 3}, out: []int{3, 2, 6}},`,
		Online: 0,
	}
}

type TestCase struct {
	In  []int `json:"in"`
	Out []int `json:"out"`
}

func NewTestCase() TestCase {
	return TestCase{
		In:  []int{1, 2, 3},
		Out: []int{3, 2, 1},
	}
}

func MJ() {
	f, err := os.OpenFile("model.json", os.O_CREATE|os.O_WRONLY, 0644)
	goutils.CheckErr(err)
	m := NewModel()
	b, err := json.MarshalIndent(m, "", "\t")
	goutils.CheckErr(err)
	f.Write(b)
}

func ToM() *Model {
	b, err := ioutil.ReadFile("model.json")
	goutils.CheckErr(err)
	var ret Model
	err = json.Unmarshal(b, &ret)
	goutils.CheckErr(err)
	return &ret
}

func MJs() {
	f, err := os.OpenFile("models.json", os.O_CREATE|os.O_WRONLY, 0644)
	goutils.CheckErr(err)
	m := NewModel()
	m2 := NewModel()

	b, err := json.MarshalIndent([]*Model{&m, &m2}, "", "\t")
	goutils.CheckErr(err)
	f.Write(b)
}

func ToMs() []Model {
	b, err := ioutil.ReadFile("models.json")
	goutils.CheckErr(err)
	var ret []Model
	err = json.Unmarshal(b, &ret)
	goutils.CheckErr(err)
	return ret
}

// http://7xku3c.com1.z0.glb.clouddn.com/models.json
func TiniuMs(_url string) []Model {
	resp, _ := http.Get(_url)
	b, err := ioutil.ReadAll(resp.Body)
	goutils.CheckErr(err)
	var ret []Model
	err = json.Unmarshal(b, &ret)
	goutils.CheckErr(err)
	return ret
}
