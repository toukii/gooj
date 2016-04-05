package model_util

import (
	"testing"
)

func TestT(t *testing.T) {
	MJ()
	m := ToM()
	t.Log(m)
}

func TestTs(t *testing.T) {
	MJs()
	m := ToMs()
	t.Log(m)
}
