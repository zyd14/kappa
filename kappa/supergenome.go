package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

type Genome struct {
    name string
    start int
    end int
    size int
}
var supergenome []Genome

func fileioerror (fname string) {
    fmt.Println ("[Fatal Error] File", fname, "cannot be accessed.")
}

func readseq (fname string) {
    fmt.Println("In development")
    supergenome = make ([]Genome,0)
    fdin,_ := os.Open (fname)
    scnr := bufio.NewScanner(fdin)
    line := ""
    start := 0
    index := ""
    for scnr.Scan() {
        line = scnr.Text()
        index = line[0: strings.Index(line, ",")]
        strnum:= line [strings.Index(line, ",") +2: len(line)]
        start,_ = strconv.Atoi(strnum)
        entry := Genome {index, start, 0, 0}
        supergenome = append (supergenome, entry)
    }
    for i := 0; i < len (supergenome) - 1; i++ {
        supergenome[i].end = supergenome[i + 1].start - 1
        supergenome[i].size = supergenome[i].end - supergenome[i].start + 1
    }
    for i := 0; i < len (supergenome) - 1; i++ {
        fmt.Printf ("SG: %d, Name: %s\n",i,supergenome[i].name)
        fmt.Printf ("\tSize: %d, Start: %d, End: %d\n", supergenome[i].size, supergenome[i].start, supergenome[i].end )
    }
}

func main () {
    argv := os.Args
    argc := len (argv)
    if argc != 2 {
        fmt.Println ("Usage: ~$ ./supergenome [super genome file name]")
        return
    } else {
        fname := argv[1]
        _,fileerr := os.Open (fname)
        if fileerr != nil {
            fileioerror (fname)
        } else {
            readseq (fname)
        }
    }
}
