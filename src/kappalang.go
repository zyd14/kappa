// Implementation of Kappa programming language
package main

import "fmt"

var input string
var online bool = true

func main () {
	for online {
		fmt.Printf (">>> ")
		size, err := fmt.Sscanln (input)
		fmt.Printf ("INPUT=[%s] SIZE=[%d]\n", input, size)
		if err != nil {
			online = false
		}
	}
}