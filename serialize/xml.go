package main

import (
	"encoding/xml"
	"fmt"
)

// 人物档案
type person struct {
	Name string `xml:"name,attr"`
	// Name string
	Age int
}

func main() {
	p1 := person{Name: "davy", Age: 18}

	var data []byte
	var err error

	// 序列化
	if data, err = xml.MarshalIndent(p1, "", " "); err != nil {
	// if data, err := xml.Marshal(p); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	p2 := new(person)
	// 反序列化
	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p2)
}