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


func TestTCR(t *testing.T){
    passingTest := NewExe("true")
    failingTest := NewExe("false")

    commit := NewExe("commit")
    revert := NewExe("revert")

    shouldCommit := tcr(passingTest, commit, revert)
    assert.Equal(t, "commit", shouldCommit.CmdLine)

    shouldRevert := tcr(failingTest, commit, revert)
    assert.Equal(t, "revert", shouldRevert.CmdLine)
}

