/* import */
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os/user"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(user.Current())
	fmt.Println(rand.Intn(10)) // 常に同じ値を返す
	fmt.Printf("%g\n", math.Sqrt(7))
}
