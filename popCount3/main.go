// Count signed bits in uint64 value by cycle
// for each of 64 bit
package main

import (
	"fmt"
	"math/rand"
)

func popCount(x uint64) int {
	var res int
	for i := uint(0); i < 63; i++ {
		res += int((x >> i) & 1)
	}
	return res
}

func main() {
	for i := 0; i < 10; i++ {
		r := rand.Uint64()
		fmt.Printf("%d -> %d\n", r, popCount(r))
	}
}
