// Echo2 prints its command line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	// Exercise 1.2: modify to print the index and value of each argument, one per line
	for i, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(i) + " " + arg
		sep = "\n"
	}
	fmt.Println(s)
}
