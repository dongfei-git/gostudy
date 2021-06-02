package main

import (
	"fmt"
	"os"
)

// 定义用户id和密码
var userId int
var userPwd string

func main() {
	// 接收用户的选择
	var key int

	// 判断是否继续显示菜单
	var loop = true

	for loop {
		fmt.Println("------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3): ")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	// 判断用户输入，显示新的提示信息
	if key == 1 {
		// 登陆
		fmt.Println("请输入用户id: ")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码: ")
		fmt.Scanf("%s\n", &userPwd)
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登陆失败")
		} else {
			fmt.Println("登陆成功")
		}
	} else if key == 2 {
		fmt.Println("用户注册逻辑")
	} else if key == 3 {
		fmt.Println("退出系统", loop)
		os.Exit(0)
	}

}
