// sqrtError project main.go
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %3.2f", float64(e))
}
func Sqrt(x float64) (float64, error) {
	var e error
	if x < 0 {
		e = ErrNegativeSqrt(x)

	}
	return math.Sqrt(x), e
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
