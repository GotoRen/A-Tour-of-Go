package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(user.Current())
}
