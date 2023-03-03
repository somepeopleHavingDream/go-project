package main

import "github.com/redis/go-redis/v9"

func singleThread() {
	// 拿到redis连接
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}

func main() {
	
}