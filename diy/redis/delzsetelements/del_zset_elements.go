package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
    host = "192.168.50.61"
    port = 6379
    db = 0
    key = "bibi_user_active_score"
)

var connection = redis.NewClient(&redis.Options{
    Addr: fmt.Sprintf("%s:%d", host, port),
    DB:   db,
})

var ctx = context.Background()

// 往redis中添加一个zset
func add_zset_to_redis() {
    // Create a slice of redis.Z with 2,000,000 elements with score 0 and 700,000 elements with score 1
    zset_data := make([]redis.Z, 2700000)
    for i := 0; i < 2700000; i++ {
        if i < 2000000 {
            zset_data[i] = redis.Z{Score: 0, Member: fmt.Sprintf("%d", i)}
        } else {
            zset_data[i] = redis.Z{Score: 1, Member: fmt.Sprintf("%d", i)}
        }
    }

    // Add the zset to Redis
    connection.ZAdd(ctx, key, zset_data...)

    // Print the number of elements in the zset
    fmt.Println(connection.ZCard(ctx, key).Val())
}

// 统计zset中的各个分值的元素的占比
func percentage_score() {
    // Use Redis command ZCOUNT to count the number of elements in the zset with score 0
    num_score_0 := connection.ZCount(ctx, key, "0", "0").Val()

    // Use Redis command ZCARD to count the total number of elements in the zset
    total_num := connection.ZCard(ctx, key).Val()
    if total_num == 0 {
        fmt.Println("total_num", total_num)
        return
    }

    // Calculate the percentage of elements with score 0
    score_0 := float64(num_score_0) / float64(total_num) * 100
    score_other := 100 - score_0
    percentage_score_0 := int(score_0 + 0.5)
    percentage_score_other := int(score_other + 0.5)

    // Print the percentage of elements with score 0 and 1
    fmt.Println("num_score_0", num_score_0)
    fmt.Println("total_num", total_num)
    fmt.Printf("percentage_score_0: %d%%\n", percentage_score_0)
    fmt.Printf("percentage_score_other: %d%%\n", percentage_score_other)
}

func main() {
    add_zset_to_redis()
    percentage_score()
}