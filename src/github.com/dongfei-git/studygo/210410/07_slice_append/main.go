package main

import "fmt"

func main() {
	s1 := []string{"beijing", "shanghai", "guangzhou"}
	// s1[3] = "shengzhen"  // 错误的写法, 索引越界
	s1 = append(s1, "shengzhen")  // 调用append函数必须用原来的切片接收返回值

	fmt.Println(s1)
}