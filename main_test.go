package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCmd(t *testing.T){
	succeed := Cmd("ls /")
	assert.Equal(t, true, succeed)
	fail := Cmd("ls /does_not_exist")
	assert.Equal(t, false, fail)
}
