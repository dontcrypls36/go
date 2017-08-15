// fibonacciClosure project main.go
package main

import (
	"fmt"
)

func fibonacci() func() int {
	first := 0
	second := 0
	third := 1
	return func() int {
		first, second, third = second, third, second+third
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
