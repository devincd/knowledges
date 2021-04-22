/*
Go has various value type including strings,integers,
floats,booleans,etc. Here are a few basic examples.

Strings,which can be added together with +.
Integers and floats.
Booleans, with boolean operators as you'd expect
*/
package main

import "fmt"

func main() {

	fmt.Println("go" + "lang")

	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

/*
$ go run values.go
golang
1+1= 2
7.0/3.0 2.3333333333333335
false
true
false
*/
