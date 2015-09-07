package gooj

import (
	"encoding/json"
	"github.com/shaalx/goutils"
	"os"
	"testing"
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
	Desc      string     `json:"desc"`
	FuncName  string     `json:"func_name"`
	ArgsType  []string   `json:"args_type"`
	RetsType  []string   `json:"rets_type"`
	TestCases []TestCase `json:"test_cases"`
}

func NewModel() Model {
	return Model{
		Desc:      "reverse the array",
		FuncName:  "reverse",
		ArgsType:  []string{"[]int"},
		RetsType:  []string{"[]int"},
		TestCases: []TestCase{NewTestCase()},
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

func TestT(t *testing.T) {
	MJ()
}
