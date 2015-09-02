package goojt

import (
	"github.com/qiniu/log"
	"testing"
	"time"
)

func Oj(result chan string) {
	retc := make(chan bool)
	t1 := time.Now()
	go func() {
		ret := {{.FUNC}}()
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

func TestOj(t *testing.T) {
	result := make(chan string)
	go Oj(result)
	log.Info(<-result)
}
