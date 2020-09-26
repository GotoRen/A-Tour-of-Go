/* reader */
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 入力元のReader ここではstringを読み取るReader
	r := strings.NewReader("Hello, Reader!")

	// Readerから読み取った結果を格納するスライス
	b := make([]byte, 8)
	for {
		// 読み取ったバイト数がn、発生したエラーがerr（エラーが発生してなければnil）、読み取ったデータは引数のスライスに格納される
		n, err := r.Read(b)
		fmt.Printf("n: %v, err: %v b: %v\n", n, err, b)
		fmt.Printf("b[:n]: %q\n", b[:n])
		// 終端まで読み取ったら、エラーとしてio.EOFを返す決まりになっている
		if err == io.EOF {
			break
		}
	}
}
