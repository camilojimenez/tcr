package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCmd(t *testing.T){
	assert.Equal(t, true, Cmd("true"))
	assert.Equal(t, false, Cmd("false"))
    assert.Equal(t, true, Cmd("ls /bin")) // string
    assert.Equal(t, false, Cmd("ls /fake")) // string
}


//func TestTCR(t *testing.T){
//    test := Exe("
//    commit := returner(true)
//    revert := returner(true)
//    assert.Equal(t, commit, tcr(test, commit, revert))
//
//
//}

