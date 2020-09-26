/* stringer */
package main

import "fmt"

/* 構造体 */
type Person struct {
	Name string
	Age  int
}
type IPAddr [4]byte

/* Stringer */
func (p *Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	// Person構造体はString()メソッドを実装しているため、fmt.Stringerインターフェースとして扱えると見做す
	a := &Person{"Ren Goto", 20}
	z := &Person{"Suzuka Tomita", 19}
	fmt.Println(a, z)

	fmt.Println()

	// IPAddr構造体はString()メソッドを実装しているため、fmt.Stringerインターフェースとして扱えると見做す
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	fmt.Println(hosts)
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
