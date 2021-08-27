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

	b := 'C'
	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%U\n", b)
	fmt.Printf("%e\n", 10.2)
	return
}
