package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

// output
// 0 -> 4
// 1 -> 4
// 2 -> 4
// 3 -> 4
// 4 -> 4

// 参考解析
// for range循环的时候会创建每个元素的副本，而不是元素的引用，并且这个副本永远都是一个，所以 m[key] = &val 取的都是变量 val 的地址，所以
// 最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3
// https://github.com/golang/go/wiki/CommonMistakes
