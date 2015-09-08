package gooj

import (
	"github.com/qiniu/log"
	"reflect"
	"testing"
	"time"
)

func Oj(result chan string) {
	retc := make(chan bool)
	t1 := time.Now()
	go func() {
		for _, it := range testcases {
			ret := FUNC(it.in)
			if !reflect.DeepEqual(ret, it.out) {
				retc <- false
				log.Println(it.in)
				log.Println("want", it.out)
				log.Println("get", ret)
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
	log.Info(time.Now().Sub(t1))
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
	go Oj(result)
	log.Info(<-result)
}

func Hack() {
	for {
		log.Print("...")
	}
}

func FUNC(in []int) []int {
	// Hack()
	leng := len(in)
	l := leng / 2
	for i := 0; i < l; i++ {
		in[i], in[leng-1-i] = in[leng-1-i], in[i]
	}
	return in
}

func TestRender(t *testing.T) {
	GenerateOjModle("./tw2", ToM())
}
