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

