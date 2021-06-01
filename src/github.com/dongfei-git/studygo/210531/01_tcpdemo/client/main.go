package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	// 从终端读取一行输入的数据
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}

		// 如果输入exit则退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}

		// 将line发送给服务器端
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write() err=", err)
		}
		// fmt.Printf("客户端发送了%d个字节的数据\n", n)
		// fmt.Println("conn suc conn=", conn)
	}
}
