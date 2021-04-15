package main

import "fmt"

func main() {
	var s1 = make([]map[int]string, 5, 10)

	s1[0] = make(map[int]string)
	s1[0][10] = "10"

	fmt.Println(s1)
}