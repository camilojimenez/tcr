package main

import (
    "os/exec"
    "strings"
)

func Cmd(cmdline string) bool {
    cmdArgs := strings.Fields(cmdline)
    cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
    if err := cmd.Run(); err == nil {
        return true
    }
    return false
}

type Exe struct {
    CmdLine string
    // TODO: add Cmder
}

func NewExe(cmdLine string) (*Exe) {
    return &Exe{
        CmdLine: cmdLine,
    }
}

func (e *Exe) Run() bool {
    return Cmd(e.CmdLine)
}

func tcr(test, commit, revert *Exe) *Exe {
    if test.Run() {
        return commit
    }
    return revert
}


func main(){
    test := NewExe("make test")
    commit := NewExe("git commit -a")
    revert := NewExe("git checkout HEAD main.go")

    tcr(test, commit, revert)
}
