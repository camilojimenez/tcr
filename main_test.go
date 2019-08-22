package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCmd(t *testing.T){
	succeed := Cmd("true")
	assert.Equal(t, true, succeed)
	fail := Cmd("false")
	assert.Equal(t, false, fail)
}


//func TestTCR(t *testing.T){
//    test := Exe("
//    commit := returner(true)
//    revert := returner(true)
//    assert.Equal(t, commit, tcr(test, commit, revert))
//
//
//}

