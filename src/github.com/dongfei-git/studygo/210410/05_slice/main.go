package main

import "fmt"

func main() {
	// 定义一个存放类型元素的切片
	var s1 []int
	var s2 []string

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"hello", "world", "golang"}

	fmt.Println(s1, s2)

	// 由数组得到切片
	a1 := [...]int{1,3,4,5,6,7}
	s3 := a1[0:4]  // 左闭右开[0,4)
	fmt.Println(s3)
}
