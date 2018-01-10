package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = ((b[i] - 'A' + 13) % 26) + 'A'
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// io.Copy(os.Stdout, s)

	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
