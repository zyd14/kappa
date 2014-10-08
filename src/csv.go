package main

import "errors"
import "os"

type csv_t struct {
	data [] [] string
}

func CsvParse (fpath string) (cdata csv_t, err error) {
	file, err := os.Open (fpath)
	if err != nil {
		return nil, nil
	} else {
		var cdata csv_t
		cdata.data = make ([] [] string, 0)
	}
}