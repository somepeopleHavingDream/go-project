package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)

	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

func sampleReadFromString() {
	data, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(data)
}

func sampleReadStdin() {
	fmt.Println("please input from stdin:")

	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(data)
}

func sampleReadFile() {
	file, _ := os.Open("main.go")
	defer file.Close()

	data, _ := ReadFrom(file, 9)
	fmt.Println(data)
}

func main() {
	// sampleReadFromString()
	// sampleReadStdin()
	sampleReadFile()
}