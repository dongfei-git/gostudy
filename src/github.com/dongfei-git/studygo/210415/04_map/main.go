package main

import "fmt"

func main() {
	// 声明和初始化
	var m1 = make(map[string]string, 10)

	m1["name"] = "tom"
	m1["age"] = "10"

	fmt.Println(m1)

	v, ok := m1["gender"]  // 判断某个key是否存在
	if !ok {
		fmt.Println("not found key")
	} else {
		fmt.Println(v)
	}

	for k, v := range m1 {
		fmt.Println(k,v)
	}

	delete(m1, "name")
	fmt.Println(m1)
}