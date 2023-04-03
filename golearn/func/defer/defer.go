package main

import (
	"fmt"
	"os"
)

func main() {
	func1()
	fmt.Println(fileSize1("./defer.go"))
	fmt.Println(fileSize2("./defer.go"))
}

// 根据文件名查询其大小
func fileSize2(filename string) int64 {
	// 根据文件名打开文件，返回文件句柄和错误
	f, err := os.Open(filename)

	// 如果打开时发生错误，返回文件大小为0
	if err != nil {
		return 0
	}

	// 延迟调用Close，此时Close不会被调用
	defer f.Close()

	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误，关闭文件并返回文件大小为0
	if err != nil {
		return 0
	}

	// 取文件大小
	size := info.Size()

	// 返回文件大小
	// defer机制触发，调用Close关闭文件
	return size
}

// 根据文件名查询其大小
func fileSize1(filename string) int64 {
	// 根据文件名打开文件，返回文件句柄和错误
	f, err := os.Open(filename)

	// 如果打开时发生错误，返回文件大小为0
	if err != nil {
		return 0
	}

	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误，关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}

	// 取文件大小
	size := info.Size()

	// 关闭文件
	f.Close()

	// 返回文件大小
	return size
}

func func1() {
	fmt.Println("defer begin")

	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入，位于栈顶，最先调用
	defer fmt.Println(3)

	fmt.Println("defer end")
}