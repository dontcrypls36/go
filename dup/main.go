// dup project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f)
		f.Close()
	}

}
func countLines(f *os.File) {
	fileCounts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		fileCounts[input.Text()]++
		if input.Text() == "." {
			break
		}
	}
	for _, n := range fileCounts {
		if n > 1 {
			fmt.Printf("%s\n", f.Name())
			break
		}
	}

}
