// Echo1 prints its command line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Exercise 1.3: compare runtimes between less efficient and more efficient versions of Echo
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
