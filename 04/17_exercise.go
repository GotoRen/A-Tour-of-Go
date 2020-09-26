/* Exercise: reader */
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	//fmt.Printf("debug: n=%v, err=%v\n", n, err)
	for i, v := range b {
		switch {
		case v >= 'A' && v < 'N', v >= 'a' && v < 'n':
			b[i] += 13
		case v >= 'N' && v <= 'Z', v >= 'n' && v <= 'z':
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	src := strings.NewReader("Lbh penpxrq gur pbqr!")
	fmt.Println("src: ", src)
	read := rot13Reader{src}
	io.Copy(os.Stdout, &read)
}
