package kshell

import "bufio"
import "fmt"
import "os"

import "kparser"

var running bool = true

func ShellInfo () {
    fmt.Println("Kappa 0.0.1 (2014)")
    fmt.Println("[GCC 4.9.1] on Ubuntu 14.04")
}

func Shell () {
    ShellInfo()
    scnr := bufio.NewReader(os.Stdin)
    for running {
        fmt.Printf (">>> ")
        line,_ := scnr.ReadString('\n')
        line = line[0:len(line) - 1]
        running = kparser.ParseInput(line)
    }
}
