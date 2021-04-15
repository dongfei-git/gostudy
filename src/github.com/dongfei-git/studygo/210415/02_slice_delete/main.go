package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5, 7}
	a1 = append(a1[:1], a1[2:]...)  // 删除index为1的元素
	fmt.Println(a1)
}