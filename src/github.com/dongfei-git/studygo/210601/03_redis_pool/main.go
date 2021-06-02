package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   // 表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 300, // 最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 从pool取出连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tom")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("r=", r)

	fmt.Println(conn)
}
