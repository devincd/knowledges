/*
A panic typically means something went unexpectedly
wrong. Mostly we use it to fail fast on errors that shouldn't
occur during normal operation, or that we aren't prepared
to handle gracefully.

Note that unlike some languages which use exceptions for
handling of many errors, in Go it is idiomatic to use error-indicating
return values wherever possible.
*/
package main

import "os"

func main() {
	// We'll use panic throughout this site to check for
	// unexpected errors. This is the only program on the site
	// designed to panic.
	panic("a problem")

	// A common use of panic is to abort if a function returns an
	// error value that we don't know how to (or want to) handle.
	// Here's an example of panicking if we get an unexpected
	// error when creating a new file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

/*
Running this program will cause it to panic, print an error
messages and goroutine traces, and exit with a non-zero status.

OUTPUT
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
       /.../panic.go:19 +0x39

Process finished with the exit code 2
*/
