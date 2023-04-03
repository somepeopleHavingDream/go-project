package main

import (
	"fmt"
	"go-project/golearn/package/encapsulation/model"
)

func main() {
	p := model.NewPerson("smith")
	p.SetAge(18)

	fmt.Println(p)
}