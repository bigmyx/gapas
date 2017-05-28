package main

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

var pass Pass

func TestSetKey(t *testing.T) {
	pass.Id = uuid.New()
	pass.Pass = "aaa"
	pass.TTL = time.Minute * 1
	SetKey(pass)
}

func Test_GetKey(t *testing.T) {
	r := GetKey(pass)
	if r != pass.Pass {
		t.Error(
			"For", r,
			"expected", pass.Pass,
			"got", r,
		)
	}
}
