package main

import "bufio"
import "errors"
import "fmt"
import "os"
//import "strconv"
import "strings"

type ref_t struct {
	name string
	seq [] int
}

type sam_t struct {
	refs [] ref_t
}

var sam sam_t

//================================ Signatures ================================//
func ParseArgs () (filein, fileout string, err error);
func SetRefCount (infile *os.File, sam sam_t) (count int)
//============================================================================//

/*
func is_header (line string) bool {
	return strings.HasPrefix(line, "@")
}

func is_ref (line string) bool {
	return strings.HasPrefix(line, "@SQ")
}

func is_entry (line string) bool {
	return !strings.HasPrefix(line, "@")
}*/

/**
 * Parse command line arguments
 */
func ParseArgs () (filein, fileout string, err error) {
	argc := len(os.Args)

	if (argc != 3) {
		return "", "", errors.New ("./samparse [input file] [output file])")
	} else {

		if !strings.HasSuffix(os.Args[1], ".sam") {
			return "", "", errors.New (os.Args[1] + "is not a valid sam file.")
		}
		if !strings.HasSuffix(os.Args[2], ".csv") {
			return "", "", errors.New (os.Args[2] + "is not a valid csv file.")
		}
		filein, errin := os.Open (os.Args[1])
		fileout, errout := os.Create (os.Args[2])

		filein.Close()
		fileout.Close ()

		if errin != nil { return "", "", errin }
		if errout != nil { return "", "", errout }
	}
	return os.Args[1], os.Args[2], nil;
}

/**
 * Read same file and set the number of references
 */
func SetRefCount (infile *os.File, sam sam_t) (count int) {
	var c int = 0
	scnr := bufio.NewScanner(infile)
	line := scnr.Text()
	fmt.Println(line)
	line = scnr.Text()
	fmt.Println(line)
	line = scnr.Text()
	fmt.Println(line)
	for strings.HasPrefix(line, "@") {
		fmt.Println(line)
		if strings.HasPrefix(line, "@SQ") {
			c = c + 1
		}
		scnr.Text()
	}
	return c
}
/*
func parse_ref (sam sam_t, line string) (name string, size int) {
	fmt.Println("Analyzing reference sequence...")
	seqnamestart := strings.Index(line, "SN:") + 3
	seqnameend := strings.Index (line[seqnamestart: len(line)], "\t") + seqnamestart
	n:=line[seqnamestart:seqnameend]
	fmt.Println(n)
	sam.refs[2].name = n
	fmt.Println("Finishing analyzing reference sequence...")
	return n, 5
}

func parse_entry (sam sam_t, line string) {
}

func populate (seq [] int, start int, end int) {
	for i:= start - 1; i <= end - 1; i++ {
		seq[i] = seq[i] + 1
	}
}

func sam_init (scanner *bufio.Scanner, sam sam_t) {
	fmt.Println("Initializing data structures...")
	var size int = GetRefCount (scanner, sam)
	fmt.Println(size)
	refcounter := scanner
	var count int = 0
	for refcounter.Scan () {
		line := refcounter.Text ()
		if ! is_header (line) {
			break;
		}
		if is_ref (line) {
			count = count + 1
		}
	}
	sam.refs = make ([] ref_t, count)

	var refsizes [] int = make ([] int, count)
	var i int = 0
	sizecounter := scanner
	for sizecounter.Scan () {
		line := sizecounter.Text ()
		if ! is_header (line) {
			break;
		}
		if is_ref (line) {
			seqsizestart := strings.Index(line, "LN:") + 3
			seqsizeend := strings.Index (line[seqsizestart: len(line)], "\t") + seqsizestart
			name:=line[seqsizestart:seqsizeend]
			num, err := strconv.Atoi(name)
			if err == nil {
				refsizes[i] = num
				i = i + 1
			}
		}
	}

	for i:=0; i < count; i++ {
		sam.refs[i].seq = make ([] int, refsizes[i])
	}
	fmt.Println("Finishing initializing data structures...")
}

func sam_scan (scanner *bufio.Scanner, sam sam_t) {
	fmt.Println("Scanning sam file...")
	for scanner.Scan () {
		line := scanner.Text ()
		if line == "." {
			break
		} else {
			if is_ref(line) {
				parse_ref (sam, line)
			} else if is_entry (line) {
				parse_entry (sam, line)
			}
		}
	}
	fmt.Println("Finishing scanning sam file...")
}

func sam_write (writer *bufio.Writer, sam sam_t) {
	fmt.Println("Generating data file...")
	for i := 0; i < len (sam.refs); i++ {
		writer.WriteString(sam.refs[i].name)
		writer.WriteString(",")
	}
	writer.Flush()
	for i := 0; i < len (sam.refs); i++ {
		for j := 0; j < len(sam.refs[i].seq); j++ {
			val := strconv.Itoa(sam.refs[i].seq[j])
			writer.WriteString(val)
			writer.WriteString(",")
		}
		writer.Flush()
	}
	fmt.Println("Finishing generating data file...")
}*/

func WriteOutput (outfile *os.File, sam sam_t) {

}

func main () {
	inpath, outpath, err := ParseArgs ()
	if (err != nil) {
		fmt.Println(err)
		return
	}
		
	fin, errin := os.Open (inpath)
	fout, errout := os.Create (outpath)

	if errin != nil || errout != nil {
		fmt.Println(errin)
		fmt.Println(errout)
		return
	}

	var ref int = SetRefCount (fin, sam)
	fmt.Println("Total Reference:", ref)
	WriteOutput(fout, sam)
	
/*
	filein, errin := os.Open (os.Args[1])
	fileout, errout := os.Create (os.Args[2])
	
	if errin == nil && errout == nil {
		scanner := bufio.NewScanner(filein)
		writer := bufio.NewWriter (fileout)
		sam_init (filein, sam)
		sam_scan (filein, sam)
		sam_write (fileout, sam)
	} else {
		fmt.Println("I/O error. Please try later.")
	}	*/
}