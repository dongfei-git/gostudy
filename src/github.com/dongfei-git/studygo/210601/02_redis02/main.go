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
	defer conn.Close()
	fmt.Println("conn suc...", conn)

	// set操作
	_, err = conn.Do("HSet", "user01", "name", "zhangsan", "age", 29)
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("set suc...")

	// get操作
	reply, err := redis.StringMap(conn.Do("HGetAll", "user01"))
	if err != nil {
		fmt.Println("conn.Do get err=", err)
		return
	}
	fmt.Println(reply)
}
