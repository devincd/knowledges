package main

import "fmt"


type User struct {
	Name string
	Age int
}

func main() {
	a := User{
		Name: "caodi",
		Age:123,
	}

	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a)
	return
}
