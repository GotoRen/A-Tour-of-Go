/* default select */
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)  // 優先度: 2
	boom := time.After(500 * time.Millisecond) // 優先度: 3
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)  // 優先度: 1
		}
	}
}
