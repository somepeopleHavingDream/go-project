package main

import (
	"bufio"
	"fmt"
	"os"
)

// exe main.go
func main() {
	if len(os.Args) < 2 {
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line int
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		if !isPrefix {
			line++
		}
	}

	fmt.Println(line)
}