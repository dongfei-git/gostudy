package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	// 循环接收客户端发送的数据
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		// fmt.Printf("服务器在等待客户端(%s)发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 如果对方没发发送则一直阻塞在此
		if err == io.EOF {
			fmt.Println("客户端已退出，Read err=", err)
			return
		}
		// 显示客户端发送内容
		fmt.Print(string(buf[:n]))
	}
}

func main() {

	fmt.Println("服务器开始监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()

	// 循环等待客户端连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc conn=%v, ip=%v\n", conn, conn.RemoteAddr().String())
		}

		go process(conn)
	}

	// fmt.Printf("listen suc=%v", listen)

}
