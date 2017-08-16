// rot13Reader project main.go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	var n int
	var err error
	n, err = rot.r.Read(b)
	for i := 0; i < len(b); i++ {
		if (b[i] >= 110 && b[i] <= 122) || (b[i] >= 78 && b[i] <= 90) {
			b[i] = b[i] - 13
		} else if (b[i] >= 97) && (b[i] <= 109) || (b[i] >= 65 && b[i] <= 77) {
			b[i] = b[i] + 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
