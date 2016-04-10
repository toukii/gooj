package utils

import (
	"fmt"
	"github.com/shaalx/goutils"
	"strings"
)

const (
	FAIL = "FAIL"
	PASS = "PASS"
)

type Result struct {
	State       string
	Pass        Pass
	Fail        Fail
	RunCostTime string
	Content     string
}

func (r *Result) String() string {
	if r.State == "PASS" {
		return fmt.Sprintf("PASS, cost: %s", r.RunCostTime)
	}
	return fmt.Sprintf("FAIL, %s, cost: %s", r.Fail, r.RunCostTime)
}

type Fail struct {
	TestCase  string
	RunResult string
	ErrorInfo string
}

func (f *Fail) String() string {
	ret := fmt.Sprintf("\tTestCase: %s\n\tRunResult: %s\n", f.TestCase, f.RunResult)
	if len(f.ErrorInfo) > 0 {
		ret += fmt.Sprintf("\tErrorInfo: %s", f.ErrorInfo)
	}
	return ret
}

type Pass struct {
}

func AnalyseBytes(input_bs []byte) *Result {
	// fmt.Println("=============================")
	result := Result{}
	input := goutils.ToString(input_bs)
	result.Content = input
	splt := strings.Split(input, "\n")
	length := len(splt)
	if length >= 5 {
		if splt[length-3] == "PASS" {
			result.Pass = AnalysePass(splt)
			result.State = "PASS"
		}
	}
	// fmt.Println(splt)
	// fmt.Println(splt[length-2])
	if strings.Contains(splt[length-2], "FAIL") {
		result.Fail = AnalyseFail(splt)
		result.State = "FAIL"
	}
	// fmt.Println(result)
	// fmt.Println("=============================")
	return &result
}

func Analyse(input string) *Result {
	// fmt.Println("=============================")
	result := Result{}
	result.Content = input
	splt := strings.Split(input, "\n")
	length := len(splt)
	if length >= 5 {
		if splt[length-3] == "PASS" {
			result.Pass = AnalysePass(splt)
			result.State = "PASS"
		}
	}
	/*for i, it := range splt {
		// fmt.Println(i, it)
	}*/
	// fmt.Println(splt[length-2])
	if strings.Contains(splt[length-2], "FAIL") {
		result.Fail = AnalyseFail(splt)
		result.State = "FAIL"
	}
	cost := strings.Split(splt[length-2], "\t")
	if len(cost[len(cost)-1]) <= 6 {
		result.RunCostTime = cost[len(cost)-1]
	}
	// fmt.Println(result)
	// fmt.Println("=============================")
	return &result
}

func AnalysePass(splt []string) Pass {
	// length := len(splt)
	return Pass{}
}

func AnalyseFail(splt []string) Fail {
	length := len(splt)
	// fmt.Println(splt[length-1])
	testcase := false
	fail := Fail{}
	for i := length - 1; i >= 0; i-- {
		if strings.Contains(splt[i], "TestCase:") {
			tc := splt[i]
			// fmt.Println(tc)
			res := strings.Split(tc, "RunResult")
			fail.TestCase = res[0]
			fail.RunResult = "RunResult" + res[1]
			testcase = true
			break
		}
	}
	if !testcase {
		// fmt.Println(splt[1:length])
		for i := 1; i < length; i++ {
			fail.ErrorInfo += splt[i] + "\n"
		}
	}
	// fmt.Println(fail)
	return fail
}
