package gooj

import (
	"encoding/json"
	"fmt"
	"github.com/shaalx/goutils"
	"io/ioutil"
	"testing"
)

func MJ() []byte {
	b, err := ioutil.ReadFile("model.json")
	goutils.CheckErr(err)
	var ret interface{}
	err = json.Unmarshal(b, &ret)
	goutils.CheckErr(err)
	fmt.Println(ret)
	return b
}

func TestT(t *testing.T) {
	MJ()
}
