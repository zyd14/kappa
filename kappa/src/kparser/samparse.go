package kparser

import "bufio"
import "errors"
import "fmt"
import "os"
import "strconv"
import "strings"

type ref_t struct {
	name string
	seq []int
}

type sam_t struct {
	refs []ref_t
}

type sam_i struct {
	size []int
	name []string
}

var seqdata sam_t
var seqinfo sam_i

/**
 * Parse integer from string that starts with an integer
 */
func FetchInt (str string) (val int) {
	size := len (str)
	var pos int = 0;
	for i:=0; i < size; i++ {
		if str[i] > '9' || str[i] < '0' {
			break
		} else {
			pos = pos + 1
		}
	}
	num, err:= strconv.Atoi(str[0: pos])
	if err != nil {
		return -1
	} else {
		return num
	}
}

/**
 * Parse each matching case of gene
 */
func GetGeneInfo (line string) (seqname string, start int, end int) {
	var pieces []string = strings.Split(line, "\t")
	name := pieces [2]
	seqstart, err := strconv.Atoi(pieces[3])
	seqend := FetchInt (pieces[5])
	if err != nil {
		return "", 0, 0
	} else {
		return name, seqstart, int(seqend) + seqstart
	}
}

/**
 * Parse an individual line that contains reference sequence information,
 * including the name of the refrence sequence and the length
 */
func GetRefInfo (line string) (name string, size int, err error) {
	fmt.Println("Analyzing reference sequence...")
	namestart := strings.Index(line, "SN:") + 3
	nameend := strings.Index (line[namestart: len(line)], "\t") + namestart
	seqname:=line[namestart:nameend]

	// CAUTION: Must make sure that LN:9999 has a tab right after. Otherwise
	// program will not execute and cause runtime error!
	sizestart := strings.Index(line, "LN:") + 3
	sizeend := strings.Index (line[sizestart: len(line)], "\t") + sizestart
	strval:=line[sizestart:sizeend]
	seqval, err:= strconv.Atoi(strval)
	if err != nil {
		return "", 0, err
	} else {
		fmt.Println("Finishing analyzing reference sequence...")
		return seqname, seqval, nil
	}
}

/**
 * Parse command line arguments and return the path for input and output files
 */
func ParseArgs () (filein, fileout string, err error) {
	fmt.Println("Analyzing command line arguments...")
	argc := len(os.Args)

	if (argc != 3) {
		return "", "", errors.New ("./samparse [input.sam] [output.csv]")
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
	fmt.Println("Finishing analyzing command line arguments...")
	return os.Args[1], os.Args[2], nil;
}

/**
 * Initialize data structure and return the number of reference scanned
 */
func SamInit (filein string, samdata * sam_t, saminfo * sam_i) (err error) {
	fmt.Println("Initializing data structures...")
	samdata.refs = make ([]ref_t, 0)
	saminfo.size = make ([]int, 0)
	saminfo.name = make ([]string, 0)
	file, errin := os.Open (filein)

	// If file cannot be open, then return error with 0 reference sequence
	if errin != nil {
		file.Close()
		fmt.Println("I/O error occured. Aborting action...")
		return errin
	} else {
		c := 0
		scnr := bufio.NewScanner(file)
		line := ""

		// This for loop can be inefficient when input is large. A potential
		// bottleneck for this program
		for scnr.Scan() {
			if strings.HasPrefix(line, "@SQ") {
				c = c + 1
				name, size, err := GetRefInfo (line)
				if (err == nil) {
					saminfo.size = append (saminfo.size, size)
					saminfo.name = append (saminfo.name, name)
				}
			}
			line = scnr.Text()
		}
		file.Close()

		for i:= 0; i < len(saminfo.size); i++ {
			ref := ref_t {name: saminfo.name[i], seq: make ([]int, saminfo.size[i])}
			samdata.refs = append(samdata.refs, ref)
		}

		fmt.Println("Finishing initialzing data structures...")
		return nil
	}
}

/**
 * Parse gene information
 */
func SamParse (filein string, samdata * sam_t, saminfo * sam_i) (err error) {
	fmt.Println ("Parsing data...")

	file, errin := os.Open (filein)

	// If file cannot be open, then return error with 0 reference sequence
	if errin != nil {
		file.Close()
		fmt.Println("I/O error occured. Aborting action...")
		return errin
	} else {

		scnr := bufio.NewScanner(file)
		line := scnr.Text()
		// This for loop can be inefficient when input is large. A potential
		// bottleneck for this program
		for scnr.Scan() {
			if !strings.HasPrefix(line, "@") && len(line) > 10{
				name, start, end := GetGeneInfo (line)
				if end > start {
					Populate (samdata, saminfo, name, start, end)
				}
			}
			line = scnr.Text()
		}
		file.Close()
	}
	fmt.Println ("Finishing parsing data...")
	return nil
}

/**
 * Add base pair count to each position
 */
func Populate (samdata * sam_t, saminfo * sam_i, name string, start int, end int) (err error) {

	if (start >= end) {
		return errors.New ("Invalid gene position: start=" + strconv.Itoa(start) + " and end=" + strconv.Itoa(end))
	}

	var pos int = -1
	for i:= 0; i < len(saminfo.name); i++ {
		if (strings.EqualFold(name, saminfo.name[i])) {
			pos = i
			break
		}
	}

	// If reference gene is not found, return right away. Future improvement should
	// consider adding corresponding message
	if pos == -1 {
		return errors.New ("Reference not found " + name)
	}

	for i:=start; i < end; i++ {
		samdata.refs[pos].seq[i] = samdata.refs[pos].seq[i] + 1
	}
	return nil
}

/**
 * Write data to specified output file
 */
func SamGenOutput (fileout string, samdata sam_t, saminfo sam_i) (err error) {
	fmt.Println("Generating output data...")
	fout, eout := os.Create (fileout)
	if eout != nil {
		return errors.New("Cannot create file: " + fileout + "\nAction aborted")
	} else {
		writer := bufio.NewWriter(fout)
		for i:= 0; i < len(saminfo.name); i++ {
			writer.WriteString (saminfo.name[i] + "\n")
			for j:= 0; j < saminfo.size[i]; j++ {
				index := strconv.Itoa(j + 1)
				val := strconv.Itoa(samdata.refs[i].seq[j])
				writer.WriteString (index + "," + val + "\n")
			}
		}
		fmt.Println("Finishing generating output data...")
		return nil
	}
}

/**
 * The main function of samparse
 *
func main () {
	inpath, outpath, err := ParseArgs ()
	if (err != nil) {
		fmt.Println(err)
		return
	}

	err = SamInit (inpath, &seqdata, &seqinfo)
	if err != nil {
		fmt.Println("File cannot open")
		return
	} else {
		fmt.Println("References:", len(seqdata.refs))
	}

	err = SamParse(inpath, &seqdata, &seqinfo)

	SamGenOutput (outpath, seqdata, seqinfo)
}
*/
