package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypot(3, 4))

	a, b := typedTwoValues()
	fmt.Println(a, b)
}

// 普通函数声明（定义）
func hypot(x, y float64) float64 {
	return math.Sqrt(x * x + y * y)
}

// 函数的返回值，同一种类型返回值
func typedTwoValues() (int, int) {
	return 1, 2
}