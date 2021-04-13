// 数组的基本操作

// 什么是数组？
// 数组（array）是一种线性表数据结构，它用一组连续的内存空间来存储一组具有相同类型的数据

// 什么是线性表？
// 线性表就是数据排成一条线一样的结构，每个线性表上的数据最多只有前和后两个方向，数组，链表，队列，栈等也是线性表
// 的结构

package main

import "fmt"

var (
	array = [5]int{1, 2, 3, 4, 5}
)

func main() {
	fmt.Println(array)

	fmt.Println(insert(array, 2, 6))

	fmt.Println(delete(array, 2))
}

func insert(arr [5]int, index, value int) [6]int {
	temp := [6]int{}
	for i := 0; i < len(temp); i++ {
		if i < index {
			temp[i] = arr[i]
		} else if i == index {
			temp[i] = value
		} else {
			// 处于 index 后面的数组元素，往右移动
			temp[i] = arr[i-1]
		}
	}
	return temp
}

func delete(arr [5]int, index int) [4]int {
	temp := [4]int{}
	for i := 0; i < len(temp); i++ {
		if i < index {
			temp[i] = arr[i]
		} else {
			// 处于 index 后面的数组元素，往左移动
			temp[i] = arr[i+1]
		}
	}
	return temp
}
