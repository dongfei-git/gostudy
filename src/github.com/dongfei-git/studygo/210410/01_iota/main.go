package main

import "fmt"

// iota当const出现的时候iota被置为0
const (
	a1 = iota
	a2 // a2不赋值与上面一样
	_
	a3
)

const (
	b1, b2 = iota + 1, iota + 2
	b3, b4 = iota + 1, iota + 2
)

const (
	_  = iota             // iota = 0
	KB = 1 << (10 * iota) // iota = 1; << 在二进制中将1向左移10位, 1 -> 10000000000 = 1024
	MB = 1 << (10 * iota) // iota = 2
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("a1=", a1)
	fmt.Println("a2=", a2)
	fmt.Println("a3=", a3)

	fmt.Println("b1=", b1)
	fmt.Println("b2=", b2)
	fmt.Println("b3=", b3)
	fmt.Println("b4=", b4)
}
