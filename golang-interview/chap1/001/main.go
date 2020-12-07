package main

import "fmt"

func main() {
	deferCall()
}

func deferCall() {
	defer func() { fmt.Println("before doing") }()
	defer func() { fmt.Println("doing") }()
	defer func() { fmt.Println("after doing") }()

	panic("trigger a panic")
}
