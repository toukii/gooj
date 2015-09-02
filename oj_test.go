package gooj

import (
	"github.com/qiniu/log"
	"testing"
	"time"
)

func Hack() {
	for {
		log.Print("...")
	}
}

func FUNC(arg ...int) []int {
	Hack()
	ret := make([]int, 3)
	ret[0], ret[1], ret[2] = 1, 2, 3
	time.Sleep(1e8)
	return ret
}

func Oj(result chan string) {
	retc := make(chan bool)
	t1 := time.Now()
	go func() {
		ret := FUNC()
		log.Println(ret)
		retc <- true
	}()

	select {
	case <-time.After(5e8):
		result <- "TIMEOUT"
	case <-retc:
		log.Info(time.Now().Sub(t1))
		result <- "AC"
	}
}

// func TestOj(t *testing.T) {
// 	result := make(chan string)
// 	go Oj(result)
// 	log.Info(<-result)
// }

func TestRender(t *testing.T) {
	content := `package goojt
	
import "time"

func makeis() []int {
ret := make([]int, 3)
ret[0], ret[1], ret[2] = 1, 2, 3
time.Sleep(1e8)
return ret
}`
	GenerateOjModle("./t2", "makeis", content)
}
