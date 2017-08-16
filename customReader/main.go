// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	var e error
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), e
}
func main() {
	reader.Validate(MyReader{})
}
