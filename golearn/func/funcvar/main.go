package main

import "fmt"

func fire() {
	fmt.Println("fire")
}

func main() {
	var f = fire
	f()
}