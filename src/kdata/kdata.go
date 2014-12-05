package kdata

type SamRef struct {
	gname string
	gsize int
}

type SamEntry struct {

}

//===================== Implementation of Stack ===============================//
type Stack struct {
	top *Element
	size int
}

type Element struct {
	value interface{}
	next *Element
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push (value interface{}) bool {
	s.top = &Element{value, s.top}
	s.size++
	return true
}

func (s *Stack) Pop () (value interface{}) {
	if s.size > 0{
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) Peek () (value interface{}) {
	return s.top.value
}

func (s *Stack) Empty () bool {
	return s.size == 0
}

//====================== Implementation of Queue =============================//
type Queue struct {
	front *Element
	size int
	back *Element
}

func (q *Queue) Add (value interface{}) bool{
	q.back = &Element{value, q.back}
	q.size++
	return true
}

func (q *Queue) Peek () (value interface{}) {
	return q.front.value
}

func (q *Queue) Remove () (value interface{}) {
	if q.size > 0 {
		value, q.front = q.front.value, q.front.next
		q.size--
		return
	}
	return nil
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
