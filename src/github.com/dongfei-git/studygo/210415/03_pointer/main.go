package main

import "fmt"

func main() {
	n := 18
	fmt.Println(&n)

	p := &n  // 取地址
	fmt.Printf("%T\n", p)

	fmt.Println(*p)  // 取值

	var a = new(int)  // new申请内存空间
	*a = 100
	fmt.Println(a, *a)
}