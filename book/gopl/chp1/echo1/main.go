package main

import (
	"fmt"
	"os"
)

// 和大多数变成语言类似，slice用于区间索引时，Go语言也采用左闭右开形式，即区间包括第一个索引元素，不包括最后一个

// Echo1 prints its command-line arguments
func main() {
	// var 申明定义了连个 string 类型的变量 s 和 sep，变量会在声明时直接初始化，如果变量没有显示初始化，则被隐式地赋予其类型的零值(zero
	// value)
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// 效率低下
// echo 程序可以每循环一次输出一个命令行参数，这个版本是不断地把新文本追加到末尾来构造字符串，字符串 s 开始为空，即值为""，每次循环会添加
// 一些文本；第一次迭代后，还会再插入一个空格，因此循环结束时每个参数中间都有一个空格，这是一次二次加工，当参数数量庞大时，开销很大
