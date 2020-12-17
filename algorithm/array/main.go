// 概念
// 数组: 数组(Array)是一种线性表结构，它用一组连续的内存空间，来存储一组具有相同类型的数据
// 线性表: 线性表就是数据排成像一条线一样的结构，每个线性表上的数据最多只有前和后两个方向，其实除了数组，链表，队列，栈等也是线性表结构
// 非线性表: 比如二叉树，堆，图等，之所以叫非线性，是因为在非线性表中，数据之间并不是简单的前后关系
// 随机访问: 由于数组是连续的内存空间和相同类型的数据，所以数组具备随机访问的能力，但是连续的内存空间和相同类型的数据这两个限制也让数组的很
// 多操作变得非常低效，比如要想在数组中删除，插入一个数据，为了保证连续性，就需要做大量的数据搬移工作。

// 通过下面的寻址公式: 计算出该元素存储的内存地址
// a[i]_address = base_address + i * data_type_size
// 根据这个公式数组可以做到随机访问

package main

import (
	"fmt"
)

func main() {
	arr := [4]int{1, 2, 4, 5}

	fmt.Println(arr)
	fmt.Println(Insert(arr, 2, 3))
	fmt.Println(Delete(arr, 2))

	// 数组变量的地址 = 数组中第一个元素的地址
	fmt.Printf("%p\n", &arr)
	fmt.Println(&arr[0])
}

// 往数组中插入一个元素
func Insert(arr [4]int, index, target int) [5]int {
	temp := [5]int{}
	for i := 0; i < len(temp); i++ {
		if i < index {
			temp[i] = arr[i]
		} else if i == index {
			temp[i] = target
		} else {
			// 处于index后面的数组元素，往右移动
			temp[i] = arr[i-1]
		}
	}
	return temp
}

// 往数组中删除一个元素
func Delete(arr [4]int, index int) [3]int {
	temp := [3]int{}
	for i := 0; i < len(temp); i++ {
		if i < index {
			temp[i] = arr[i]
		} else {
			// 处于index后面的数组元素，往左移动
			temp[i] = arr[i+1]
		}
	}
	return temp
}
