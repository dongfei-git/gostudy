package main

import "fmt"

// 跳出多层for循环
func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto FLAGLABEL // 跳到指定的标签
			}
			fmt.Printf("%d-%c\n", i, j)
		}
	}

FLAGLABEL:
	fmt.Println("over")
}
