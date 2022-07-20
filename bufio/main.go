package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	strReader := strings.NewReader("hello world")
	bufReader := bufio.NewReader(strReader)

	data, _ := bufReader.Peek(5)
	fmt.Println(data, bufReader.Buffered())

	str, _ := bufReader.ReadString(' ')
	fmt.Println(str,bufReader.Buffered())

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello ")
	fmt.Fprint(w, "world!")
	w.Flush()
}