package main

import "fmt"

// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]int
	var a2 [4]bool

	fmt.Printf("a1: %T, a2: %T\n", a1, a2)
}