/* 可変長map */
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	//a := Vertex{40.68433, -74.39967}
	//fmt.Println(a.Lat)

	m = make(map[string]Vertex)
	fmt.Println(m)			    // map[]
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["hoge"] = Vertex{
		12.345, -67.890,
	}
	fmt.Println(m)				// map[Bell Labs:{40.68433 -74.39967} hoge:{12.345 -67.89}]
	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}

}
