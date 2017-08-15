// Count signed bits in uint64 value with 'x & (x - 1)' expression
package main

import (
	"fmt"
	"math/rand"
)

func popCount(x uint64) int {
	var res int
	for x != 0 {
		x = x & (x - 1)
		res++
	}
	return res
}

func main() {
	for i := 0; i < 10; i++ {
		r := rand.Uint64()
		fmt.Printf("%d -> %d\n", r, popCount(r))
	}

}
