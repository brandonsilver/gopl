// Dup2 prints the count and text of lines that appear more than once in the input. It
// reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Exercise 1.4: modify dup2 to print the names of all files in which the
	// duplicated line occurs.
	counts := make(map[string]int)
	lineAppearsInFiles := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineAppearsInFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineAppearsInFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s| appears in: ", n, line)
			sep := ""
			for fileName, _ := range lineAppearsInFiles[line] {
				fmt.Printf("%s%s", sep, fileName)
				sep = ","
			}
			fmt.Print("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineAppearsInFiles map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileNames, ok := lineAppearsInFiles[input.Text()]
		if !ok {
			fileNames = make(map[string]bool)
			lineAppearsInFiles[input.Text()] = fileNames
		}
		lineAppearsInFiles[input.Text()][f.Name()] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
