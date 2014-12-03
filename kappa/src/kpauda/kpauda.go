package main

import "bufio"
import "fmt"
import "os"
import "strings"

func GetHitName (line string) string {
    return line[1: len(line)]
}

func GetQueryName (line string) string {
    return line[7: len(line)]
}

func GetExpVal (line string) string {
    pos := strings.Index(line, "Expect")
    return line[pos + 9: len(line)]
}

func IsExpect (line string) bool {
    return strings.Contains(line, "Expect")
}

func IsQuery (line string) bool {
    return strings.Index(line, "Query=") == 0
}

func IsHit (line string) bool {
    return strings.Index(line, ">") == 0
}

func ParseArgs () (string, string) {
    argc := len(os.Args)
    if argc != 3 {
        fmt.Println ("Usage: ./kpauda [input file] [outputfile]")
        os.Exit(1)
    }
    _,errin := os.Open(os.Args[1])
    _, errout := os.Create(os.Args[2])
    if errin != nil || errout != nil {
        if errin != nil {
            fmt.Println("Cannot open file:", os.Args[1])
        }
        if errout != nil {
            fmt.Println("Cannot create file:", os.Args[2])
        }
        os.Exit(1)
    }
    os.Remove(os.Args[2])
    return os.Args[1], os.Args[2]
}

func Parse (fin, fout string) {
    fdin,_ := os.Open(fin)
    fdout,_ := os.Create(fout)
    scnr := bufio.NewScanner(fdin)
    prwr := bufio.NewWriter(fdout)
    currQuery := ""
    currHit := ""
    currExp := ""
    output := ""
    i := 0
    for scnr.Scan() {
        i++
        line := scnr.Text()
        if IsQuery(line) {
            currQuery = GetQueryName(line)
        }
        if IsHit(line) {
            currHit = GetHitName(line)
        }
        if IsExpect(line) {
            currExp = GetExpVal(line)
            output = currQuery + "\t" + currHit + "\t" + currExp + "\n"
            prwr.WriteString(output)
        }
        if i % 100000 == 0 {
            fmt.Printf("%d lines processed.\n", i)
        }
    }
}

func main () {
    fin,fout, := ParseArgs()
    Parse(fin, fout)
}
