package main

import "fmt"
import "./sam"

func main () {
	array := {1,2,3,4,5,6,7,8}
	mean := sam.Average(&array)
	fmt.Printf("Average = %f\n", mean)
}