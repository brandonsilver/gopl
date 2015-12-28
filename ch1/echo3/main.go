// Echo3 prints its command line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Exercise 1.1: modify to also print the program's name
	// Exercise 1.3: compare runtimes between less, more efficient versions of Echo
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
