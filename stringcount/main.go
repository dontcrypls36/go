// stringcount
// project count same strings in input string

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	for _, str := range strings.Fields(s) {
		m[str]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
