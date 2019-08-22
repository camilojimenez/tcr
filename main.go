package main

import (
	//"os/exec"
	"fmt"
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
	fmt.Println("hello world")
}

//func runTest(){
//
//  go test .
//}
//  git commit -a
//  git checkout HEAD main.go
//
//  inotifywait -r -e modify . >&2
//  do_test && do_commit || do_revert
// cbTerm := exec.Command("/bin/stty", cbreak, echo)    
// cbTerm.Stdin = os.Stdin                              
// err := cbTerm.Run()                                  

