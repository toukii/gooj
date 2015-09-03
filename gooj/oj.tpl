package goojt

import (
	"github.com/qiniu/log"
	"testing"
	"time"
	"reflect"
)

func Oj(result chan string) {
	retc := make(chan bool)
	t1 := time.Now()
	go func() {
		for _, it := range testcases {
			ret := {{.FUNC}}(it.in)
			if !reflect.DeepEqual(ret, it.out) {
				retc <- false
				log.Println(it.in)
				log.Println("want", it.out)
				log.Println("get", ret)
				return
			}
		}
		retc <- true
	}()

	select {
	case <-time.After((time.Duration)(len(testcases) * 5e9)):
		result <- "TIMEOUT"
	case ok := <-retc:
		log.Info(time.Now().Sub(t1))
		if ok {
			result <- "AC"
		} else {
			result <- "WRONG"
		}
	}
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
	go Oj(result)
	log.Info(<-result)
}
