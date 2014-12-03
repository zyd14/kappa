/**
* kappa bioinformatics programming environment
* @author Yubing Hou
* @date October 16, 2014
*/

package main

import "bufio"
import "fmt"
import "os"
import "strings"


type kvar_t struct {
    kvar int
}
var varlist map[string] kvar_t
var running bool = true

func sysinfo () {
    dir,_ := os.Getwd()
    fmt.Println("kappa bioinfomatics suite")
    fmt.Println("License: General Public License v3")
    fmt.Println("Current working directory:",dir)
}

func runparse (input string) bool {
    if strings.EqualFold (input, "exit") {
        fmt.Println("Bye!")
        return false
    } else {
        if strings.EqualFold (input, "samread") {
            fmt.Println("Read a sam file. In development")
            return true;
        } else if strings.EqualFold (input, "sgenomeread"){
            fmt.Println("Read a super genome file. In development")
        } else if strings.EqualFold (input, "pwd") {
            dir,_ := os.Getwd()
            fmt.Println(dir)
        }else {
            return true
        }
    }
    return true
}

func main () {
    sysinfo()
    _ = varlist
    reader := bufio.NewReader(os.Stdin)
    line := ""
    for running {
        fmt.Printf(">>> ")
        line,_ = reader.ReadString('\n')
        running = runparse (line[0: len(line) - 1])
    }
}
