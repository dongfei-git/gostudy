package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	fmt.Println("conn suc...", conn)
}
