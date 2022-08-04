package algo

import "math/rand"

const Min = int64(1)

// SimpleRand
// 简单随机算法
// 红包的数量，红包金额
// 金额单位为分，1元钱=100分
func SimpleRand(count, amount int64) int64 {
	// 当红包数量剩余一个的时候，就直接返回剩余金额
	if count == 1 {
		return amount
	}

	// 计算最大可调度金额
	max := amount - Min * count
	x := rand.Int63n(max) + Min
	return x
}