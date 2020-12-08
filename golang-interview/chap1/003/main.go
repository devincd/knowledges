package main

import "fmt"

func main() {
	example1()
	example2()
}

// case-1
func example1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func example2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
}

// case-1解析
// output:
// [0 0 0 0 0 1 2 3]
// [1 2 3 4]
// make()函数初始化slice的标准解释
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length. For example, make([]int, 0, 10) allocates an underlying array
//	of size 10 and returns a slice of length 0 and capacity 10 that is
//	backed by this underlying array.

// case-2
//func funcMui(x, y int) (sum int, error) {
//	return x+y, nil
//}

// case-2解析
// build with a error: syntax error: mixed named and unnamed function parameters
// 在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。

// case-3和解析
// new() 与 make() 的区别?
// new(T) 和 make(T, args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。
//
// new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话说就是返回一个指针，该指针指向新分配的，类型
// 为 T 的零值。适用于值类型，如整型，数组，结构体等
//
// make(T, args) 返回初始化之后的 T 类型的值，这个值并不是T类型的零值，也不是指针 *T ，是经过初始化之后的 T 类型的引用。make() 只适用
// slice，map 和 channel
