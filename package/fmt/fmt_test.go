package main

import (
	"fmt"
	"testing"
)

// $ go test -v fmt_test.go -test.run Test_Errorf
func Test_Errorf(t *testing.T) {
	const name, id = "'bueller'", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}

// $ go test -v fmt_test.go -test.run Test_Println
func Test_Println(t *testing.T) {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old.")

	// It is conventional not to worry about any
	// error returned by Println.
}

func Test_Printf(t *testing.T) {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.
}
