package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func findKeys() {
    r := redis.NewClient(&redis.Options{
        Addr:     "x",
        Password: "x",
        DB:       0,
    })

    iter := r.Scan(ctx, 0, "bibi:ktv_song_zset*", 100).Iterator()
    for iter.Next(ctx) {
        fmt.Println(iter.Val())
    }
    if err := iter.Err(); err != nil {
        panic(err)
    }
}

func main() {
    findKeys()
}
