package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13 (b byte) byte {
	switch {
	case b >= 'A' && b <= 'M':
		return (b-'A'+'N')
	case b >= 'N' && b <= 'Z':
		return (b-'N'+'A')
	case b >= 'a' && b <= 'm':
		return (b-'a'+'n')
	case b >= 'n' && b <= 'z':
		return (b-'n'+'a')
	default:
		return b
	}
}

func (rot13r *rot13Reader) Read (b []byte) (n int, err error) {
	input := make([]byte, 1)
	n, err = rot13r.r.Read(input)
	if err == io.EOF {
		return
	}
	b[0] = rot13(input[0])
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
