package main

import "./kstat"
import "fmt"

func main () {
	var data [65] int
	avg := kstat.average(data)
	fmt.Println("Average = ", avg)
}