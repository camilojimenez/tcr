package main

import (
    "fmt"
    "os/exec"
    "strings"
    "os"
)

func Cmd(cmdline string) bool {
    cmdArgs := strings.Fields(cmdline)
    cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
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
    fmt.Println(e.CmdLine)
    return Cmd(e.CmdLine)
}

func tcr(test, commit, revert *Exe) *Exe {
    if test.Run() {
        return commit
    }
    return revert
}


func main(){
    build := NewExe("make build")
    test := NewExe("make test")
    commit := NewExe("git commit -a")
    revert := NewExe("git checkout HEAD main.go")
    wait := NewExe("inotifywait -r -e modify .")

    for {
        build.Run()
        res := tcr(test, commit, revert)
        res.Run()
        wait.Run()
    }
}
