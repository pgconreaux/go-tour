package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rr.r.Read(b)
	for i := 0; i < len(b); i++ {
		// A-Z
		if b[i] >= 65 && b[i] <= 90 {
			b[i] += 13
			if b[i] > 90 {
				b[i] = b[i] - 90 + 64
			}
		}
		// a-z
		if b[i] >= 97 && b[i] <= 122 {
			b[i] += 13
			if b[i] > 122 {
				b[i] = b[i]-122+96
			}
		}
	}
	return n, err

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

