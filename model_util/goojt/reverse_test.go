package goojt

import (
	"testing"
	"time"
	"reflect"
	"fmt"
)

func Oj(result chan string, t *testing.T) {
	retc := make(chan bool)
	t1 := time.Now()
	go func() {
		for _, it := range testcases {
			ret := reverse(it.in)
			if !reflect.DeepEqual(ret, it.out) {
				t.Error("FAILED")
				fmt.Printf("TestCase:%#v, RunResult:%#v\n", it.in, ret)
				fmt.Printf("ResultWanted:%#v\n", it.out)
				retc <- false
				return
			}
			retc <- true
		}
	}()
	length := len(testcases)
	ticker := time.NewTicker(5e8)
	for i := 0; i < length; i++ {
		select {
		case <-ticker.C:
			result <- "TIMEOUT"
		case ok := <-retc:
			if !ok {
				result <- "WA"
				return
			}
		}
	}
	println(time.Now().Sub(t1))
	result <- "AC"
}

type TestCases struct {
	in  []int
	out []int
}

var testcases []TestCases

func init() {
	testcases = []TestCases{
		{in: []int{1, 2, 3}, out: []int{3, 2, 1}},
		{in: []int{1, 2, 4}, out: []int{4, 2, 1}},
		{in: []int{1, 5, 3}, out: []int{3, 5, 1}},
		{in: []int{6, 2, 3}, out: []int{3, 2, 6}},
	}
}

func TestOj(t *testing.T) {
	result := make(chan string)
	go Oj(result, t)
	<-result
}
