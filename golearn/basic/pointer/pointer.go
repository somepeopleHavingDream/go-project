package main

import (
	"flag"
	"fmt"
)

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main() {
	cat := 1
	str := "banana"
	fmt.Printf("%p %p \n", &cat, &str)

	// 使用指针变量获取命令行的输入信息
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)

	// 创建指针的另一种方法——new()函数
	newStr := new(string)
	*newStr = "Go语言教程"
	fmt.Println(*newStr)
}