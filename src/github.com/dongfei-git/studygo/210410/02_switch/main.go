package main

import "fmt"

func SwitchDemo(n int) {
	switch n {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	default:
		fmt.Println("星期七")
	}
}

func main() {
	SwitchDemo(6)
}
