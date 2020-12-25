package main

import (
	"fmt"
	"os"
)

// 和大多数变成语言类似，slice用于区间索引时，Go语言也采用左闭右开形式，即区间包括第一个索引元素，不包括最后一个

// Echo1 prints its command-line arguments
func main() {
	// var 申明定义了连个 string 类型的变量 s 和 sep，变量会在声明时直接初始化，如果变量没有显示初始化，则被隐式地赋予其类型的零值(zero value)
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
