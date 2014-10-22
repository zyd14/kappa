package kdata

type SamRef struct {
	gname string
	gsize int
}

type SamEntry struct {
	
}

/*
import "errors"
import "os"
import "strings"

type csv_t struct {
	data [] [] string
}

type tsv_t struct {
	data [] [] string
}

/**
 * Parse an input .csv file into a csv data structure
 *
func CsvParse (fpath string) (csv_t, error) {
	if !strings.HasSuffix(fpath, ".csv") {
		return nil, errors.New(fpath + " is not a valid .csv file")
	}

	_, err := os.Open (fpath)
	if err != nil {
		return nil, nil
	} else {
		var cdata csv_t
		cdata.data = make ([] [] string, 0)
		return cdata, nil
	}
}

/**
 * Parse an input .tsv file into a tsv data structure
 *
func TsvParse (fpath string) (csv_t, error) {
	return nil, nil
}*/
