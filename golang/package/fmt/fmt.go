package main

import (
	"fmt"
)

var (
	name string
)

type Human struct {
	name string
	age  int
}

func main() {
	name = "caodi"

	human := Human{
		name:"caodi",
		age:1,
	}
	fmt.Println(name, human)
	//fmt.Printf("human: %v\n", human)
	//fmt.Printf("human: %+v\n", human)
	//
	//fmt.Printf("human: %#v\n", human)
	//fmt.Printf("human: %T\n", human)
	//
	//fmt.Printf("文字百分号: %%\n")
	//fmt.Println(human)
}
