/* map */
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m := map[string]string{
		"go": "golang",
		"rb": "ruby",
		"js": "javascript",
	}
	fmt.Println(m) // map[go:golang js:javascript rb:ruby]
}
