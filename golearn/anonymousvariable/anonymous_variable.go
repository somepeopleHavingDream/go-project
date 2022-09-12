package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}

func main() {
	a, _ := GetData()
	_, b := GetData()

	fmt.Println(a, b)
}