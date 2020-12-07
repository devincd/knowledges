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

// 参考解析: defer的执行顺序是后进先出。当出现panic语句的时候，会先按照defer的后进先出的顺序执行，最后才会执行panic
