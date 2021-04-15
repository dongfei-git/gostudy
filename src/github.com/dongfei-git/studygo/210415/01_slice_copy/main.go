package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 // 赋值
	var a3 = make([]int, 3, 3)
	copy(a3, a1) // 拷贝

	fmt.Println(a1, a2, a3)

	a1[0] = 100
	fmt.Println(a1, a2, a3)
}