/*
A goroutine is a lightweight thread of execution.
*/
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Suppose we have a function call f(s). Here's how we'd call
	// that in the usual way, running it synchronously.
	f("direct")

	// To invoke this function in a goroutine, use go f(s). This
	// new goroutine will execute concurrently with th calling
	// one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous function
	// call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously
	// in separate goroutines now. Wait for them to finish(for a
	// more robust approach, use a WaitGroup).
	time.Sleep(time.Second)
	fmt.Println("done")
}

/*
OUTPUT
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
going
goroutine : 0
goroutine : 1
goroutine : 2
done
*/
